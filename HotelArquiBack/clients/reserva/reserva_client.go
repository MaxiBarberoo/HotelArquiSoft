package clients

import (
	"HotelArquiSoft/HotelArquiBack/model"

	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetReservaById(id int) model.Reserva {
	var reserva model.Reserva

	Db.Where("id = ?", id).First(&reserva)
	log.Debug("Reserva: ", reserva)
	return reserva
}

func GetReservas() model.Reservas {
	var reservas model.Reservas
	Db.Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)
	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva Created: ", reserva.ID)

	return reserva
}

func GetRooms(fecha time.Time, reserva model.Reserva) int {
	var count int

	err := Db.Select("COUNT(reservas.id)").
		Joins("JOIN hotels ON reservas.hotel_id = ?", reserva.HotelId).
		Where("? >= reservas.fecha_in AND ? <= reservas.fecha_out", fecha).
		Count(&count).Error

	if err != nil {
		log.Fatal(err)
	}

	return count
}

func GetReservasByUser(userId int) model.Reservas {
	var reservas model.Reservas
	Db.Where("UserId = ?", userId).Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}
