package imagenes

import (
	"HotelArquiSoft/HotelArquiBack/dto"
	jwtToken "HotelArquiSoft/HotelArquiBack/jwt"
	service "HotelArquiSoft/HotelArquiBack/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})
		return
	}

	err = mapstructure.Decode(claims, &imageDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}

	imageDTO, er := service.ImagenService.InsertImagen(imageDTO)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	signedtoken, err1 := jwtToken.GenerateImageToken(imageDTO)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error al obtener los datos",
		})

		return
	}
	c.JSON(http.StatusCreated, signedtoken)

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

	var tokens []string

	for _, imagen := range imagenes {
		token, err := jwtToken.GenerateImageToken(imagen)

		if err != nil {
			return
		}

		tokens = append(tokens, token)
	}

	c.JSON(http.StatusOK, tokens)

}
