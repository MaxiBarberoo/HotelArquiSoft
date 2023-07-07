package imagenes

import (
	"HotelArquiSoft/HotelArquiBack/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertImagen(imagen model.Imagen) model.Imagen {
	result := Db.Create(&imagen)

	if result.Error != nil {

		//TODO Manage Errors

		log.Error("")
	}
	log.Debug("Image Created: ", imagen.ID)
	return imagen
}

func GetImagenesByHotel(hotel_id int) model.Imagenes {
	var imagenes model.Imagenes
	Db.Where("hotel_id = ?", hotel_id).Find(&imagenes)
	log.Debug("Imagenes: ", imagenes)
	return imagenes
}
