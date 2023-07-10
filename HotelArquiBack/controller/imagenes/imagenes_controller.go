package imagenes

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	service "HotelArquiSoft/HotelArquiBack/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InsertImagen(c *gin.Context) {
	// Obtener los datos de la imagen desde el cuerpo de la solicitud
	var imageDTO dto.ImageDTO
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
	err = c.BindJSON(&imageDTO)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	imageDTO, er := service.ImagenService.InsertImagen(imageDTO)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, imageDTO)

}
func GetImagenesByHotel(c *gin.Context) {
	log.Debug("Image id to load: " + c.Param("image_id"))

	hotelId, _ := strconv.Atoi(c.Param("hotel_id"))

	var imagenes dto.ImagenesDTO
	imagenes, err := service.ImagenService.GetImagenesByHotel(hotelId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagenes)

}
