package amenitiehotel

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	jwtToken "HotelArquiSoft/HotelArquiBack/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func AssignAmenitieToHotel(c *gin.Context) {
	var amenitiehotelDto dto.AmenitieHotelDto

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

	err = mapstructure.Decode(claims, &amenitiehotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitiehotelDto, err = service.AmenitieHotelService.AssignAmenitieToHotel(amenitiehotelDto)

	signedtoken, err1 := jwtToken.GenerateAmentieHotelToken(amenitiehotelDto)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}
	c.JSON(http.StatusOK, signedtoken)
}

func SearchAmenitiesByHotel(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotel_id"))

	hotelId, _ := strconv.Atoi(c.Param("hotel_id"))

	var amenitiesHotels dto.AmenitiesHotels
	amenitiesHotels, err := service.AmenitieHotelService.SearchAmenitiesByHotel(hotelId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var tokens []string

	for _, amenitieHotel := range amenitiesHotels {
		token, err := jwtToken.GenerateAmentieHotelToken(amenitieHotel)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	c.JSON(http.StatusOK, tokens)

}
