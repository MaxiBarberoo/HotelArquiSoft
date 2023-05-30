package clients

import (
	"HotelArquiSoft/model"

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
