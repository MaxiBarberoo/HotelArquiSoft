package amenitiehotel

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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

	err = c.BindJSON(&amenitiehotelDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitiehotelDto, err = service.AmenitieHotelService.AssignAmenitieToHotel(amenitiehotelDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, amenitiehotelDto)
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

	c.JSON(http.StatusOK, amenitiesHotels)
}
