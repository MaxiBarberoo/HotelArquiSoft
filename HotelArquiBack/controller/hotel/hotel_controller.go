package hotel

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	jwtToken "HotelArquiSoft/HotelArquiBack/jwt"
	service "HotelArquiSoft/HotelArquiBack/services"

	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	token, err1 := jwtToken.GenerateHotelToken(hotelDto)

	if err1 != nil {
		return
	}

	c.JSON(http.StatusOK, token)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var tokens []string

	for _, hotel := range hotelsDto {
		token, err := jwtToken.GenerateHotelToken(hotel)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	c.JSON(http.StatusOK, tokens)
}

func HotelInsert(c *gin.Context) {
	var hotelDto dto.HotelDto

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token no proporcionado",
		})
		return
	}

	secret := "secreto"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalido",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})
		return
	}

	err = mapstructure.Decode(claims, &hotelDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	signedtoken, err1 := jwtToken.GenerateHotelToken(hotelDto)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}
	c.JSON(http.StatusCreated, signedtoken)
}
