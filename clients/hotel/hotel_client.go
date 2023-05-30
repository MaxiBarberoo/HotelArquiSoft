package hotel

import (
	"HotelArquiSoft/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)

	Db.Where("id=?", id).First(&hotel)

	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetUsers() model.Hotel {
	var hotel model.Hotel
	Db.Find(&hotel)
	log.Debug("Hotel: ", hotel)
	return hotel

}

func InsertUser(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {

		//TODO Manage Errors

		log.Error("")
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}
