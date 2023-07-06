package controller

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"net/http"
	"strconv"

	jwtReserva "HotelArquiSoft/HotelArquiBack/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

	token, err1 := jwtReserva.GenerateReservaToken(reservaDto)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No se pudo generar la token",
		})
		return
	}

	c.JSON(http.StatusOK, token)
}

func GetReservas(c *gin.Context) {
	var reservasDto dto.ReservasDto
	reservasDto, err := service.ReservaService.GetReservas()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var tokens []string

	for _, reserva := range reservasDto {
		token, err := jwtReserva.GenerateReservaToken(reserva)
		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	c.JSON(http.StatusOK, tokens)
}

func ReservaInsert(c *gin.Context) {
	var reservaDto dto.ReservaDto

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

	err = mapstructure.Decode(claims, &reservaDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}

	reservaDto, er := service.ReservaService.InsertReserva(reservaDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func GetRooms(c *gin.Context) {

	var reservaDto dto.ReservaDto

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

	err = mapstructure.Decode(claims, &reservaDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}

	if service.ReservaService.GetRooms(reservaDto) {
		c.JSON(http.StatusAccepted, gin.H{
			"disponibilidad": "true",
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"disponibilidad": "false",
		})
	}
}

func GetReservasByUser(c *gin.Context) {
	var tokens []string

	userId, _ := strconv.Atoi(c.Param("user_id"))

	var reservasDto dto.ReservasDto
	reservasDto, err := service.ReservaService.GetReservasByUser(userId)

	for _, reserva := range reservasDto {
		token, err := jwtReserva.GenerateReservaToken(reserva)
		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func GetReservasByFecha(c *gin.Context) {

	var reservaDto dto.ReservaDto

	err := c.BindJSON(&reservaDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservasDto dto.ReservasDto

	reservasDto, err = service.ReservaService.GetReservasByFecha(reservaDto)

	var tokens []string

	for _, reserva := range reservasDto {
		token, err := jwtReserva.GenerateReservaToken(reserva)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func GetHotelsByFecha(c *gin.Context) {
	var reservaDto dto.ReservaDto

	err := c.BindJSON(&reservaDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var reservasDto dto.ReservasDto
	reservasDto, err = service.ReservaService.GetHotelsByFecha(reservaDto)

	var tokens []string

	for _, reserva := range reservasDto {
		token, err := jwtReserva.GenerateReservaToken(reserva)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, tokens)

}
