package amenitie

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	jwtToken "HotelArquiSoft/HotelArquiBack/jwt"
	service "HotelArquiSoft/HotelArquiBack/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var amenitieDto dto.AmenitieDto

	amenitieDto, err := service.AmenitieService.GetAmenitieById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	token, err1 := jwtToken.GenerateAmenitieToken(amenitieDto)

	if err1 != nil {
		return
	}

	c.JSON(http.StatusOK, token)
}

func GetAmenities(c *gin.Context) {
	var amenitiesDto dto.AmenitiesDto
	amenitiesDto, err := service.AmenitieService.GetAmenities()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var tokens []string

	for _, amenitie := range amenitiesDto {
		token, err := jwtToken.GenerateAmenitieToken(amenitie)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	c.JSON(http.StatusOK, tokens)
}
