package amenitie

import (
	"HotelArquiSoft/HotelArquiBack/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie

	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("Amenitie: ", amenitie)

	return amenitie
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities
	Db.Find(&amenities)
	log.Debug("Amenities: ", amenities)
	return amenities
}
