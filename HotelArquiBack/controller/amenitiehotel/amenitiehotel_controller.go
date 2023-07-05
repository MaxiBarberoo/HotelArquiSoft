package amenitiehotel

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func AssignAmenitieToHotel(c *gin.Context) {
	var amenitiehotelDto dto.AmenitieHotelDto
	err := c.BindJSON(&amenitiehotelDto)

	// Error Parsing json param
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
