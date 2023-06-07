package controller

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetReservaById(c *gin.Context) {
	log.Debug("Reserva id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var reservaDto dto.ReservaDto

	reservaDto, err := service.ReservaService.GetReservaById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, reservaDto)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := service.ReservaService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, reservasDto)
}

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	reservaDto, er := service.ReservaService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, reservaDto)
}

func GetRooms(c *gin.Context) {

	var reservaDto dto.ReservaDto
	err := c.BindJSON(&reservaDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if service.ReservaService.GetRooms(reservaDto) {
		c.JSON(http.StatusAccepted, gin.H{
			"disponibilidad": "true",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"disponibilidad": "false",
	})
}
