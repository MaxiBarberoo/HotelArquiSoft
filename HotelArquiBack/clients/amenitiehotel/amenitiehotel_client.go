package amenitiehotel

import (
	"HotelArquiSoft/HotelArquiBack/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func AssignAmenitieToHotel(amenitieHotel model.AmenitieHotel) model.AmenitieHotel {
	result := Db.Create(&amenitieHotel)

	if result.Error != nil {

		//TODO Manage Errors

		log.Error("")
	}
	log.Debug("Amenitie asignada a hotel: ", amenitieHotel.ID)
	return amenitieHotel
}

func SearchAmenitiesByHotel(hotelId int) model.AmenitiesHotels {
	var amenitiesHotels model.AmenitiesHotels
	Db.Where("hotel_id = ?", hotelId).Find(&amenitiesHotels)
	log.Debug("Hoteles: ", amenitiesHotels)
	return amenitiesHotels
}
