package clients

import (
	"HotelArquiSoft/HotelArquiBack/model"

	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type reservaClient struct{}

type ReservaClientInterface interface {
	GetReservaById(id int) model.Reserva
	GetReservas() model.Reservas
	InsertReserva(reserva model.Reserva) model.Reserva
	GetRooms(fecha time.Time, reserva model.Reserva) int
	GetReservasByUser(userId int) model.Reservas
	GetReservasByFecha(reserva model.Reserva) model.Reservas
	GetReservasByHotel(hotelId int) model.Reservas
	GetReservasByHotelAndFecha(reserva model.Reserva) model.Reservas
	GetReservasByHotelAndUser(reserva model.Reserva) model.Reservas
	GetReservasByFechaAndUser(reserva model.Reserva) model.Reservas
	GetReservasByHotelFechaAndUser(reserva model.Reserva) model.Reservas
}

var (
	ReservaClient ReservaClientInterface
)

func init() {
	ReservaClient = &reservaClient{}
}
func (c *reservaClient) GetReservaById(id int) model.Reserva {
	var reserva model.Reserva

	Db.Where("id = ?", id).First(&reserva)
	log.Debug("Reserva: ", reserva)
	return reserva
}

func (c *reservaClient) GetReservas() model.Reservas {
	var reservas model.Reservas
	Db.Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func (c *reservaClient) InsertReserva(reserva model.Reserva) model.Reserva {
	result := Db.Create(&reserva)
	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva Created: ", reserva.ID)

	return reserva
}

func (c *reservaClient) GetRooms(fecha time.Time, reserva model.Reserva) int {
	var count int

	err := Db.Table("reservas").
		Select("COUNT(reservas.id)").
		Joins("JOIN hotels ON reservas.hotel_id = hotels.id").
		Where("? >= reservas.fecha_in AND ? <= reservas.fecha_out AND ? = hotels.id", fecha, fecha, reserva.HotelId).
		Count(&count).Error

	if err != nil {
		log.Fatal(err)
	}

	return count
}

func (c *reservaClient) GetReservasByUser(userId int) model.Reservas {
	var reservas model.Reservas
	Db.Where("user_id = ?", userId).Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func (c *reservaClient) GetReservasByFecha(reserva model.Reserva) model.Reservas {
	var reservas model.Reservas
	err := Db.Where("fecha_in >= ? AND fecha_out <= ?", reserva.FechaIn, reserva.FechaOut).Find(&reservas).Error
	if err != nil {
		return nil
	}
	return reservas
}

func (c *reservaClient) GetReservasByHotel(hotelId int) model.Reservas {
	var reservas model.Reservas
	Db.Where("hotel_id = ?", hotelId).Find(&reservas)
	log.Debug("Reservas: ", reservas)
	return reservas
}

func (c *reservaClient) GetReservasByHotelAndFecha(reserva model.Reserva) model.Reservas {
	var reservas model.Reservas
	err := Db.Where("fecha_in >= ? AND fecha_out <= ? AND hotel_id = ?", reserva.FechaIn, reserva.FechaOut, reserva.HotelId).Find(&reservas).Error
	if err != nil {
		return nil
	}
	return reservas
}

func (c *reservaClient) GetReservasByHotelAndUser(reserva model.Reserva) model.Reservas {
	var reservas model.Reservas
	err := Db.Where("hotel_id = ? AND user_id = ?", reserva.HotelId, reserva.UserId).Find(&reservas).Error
	if err != nil {
		return nil
	}
	return reservas
}

func (c *reservaClient) GetReservasByFechaAndUser(reserva model.Reserva) model.Reservas {
	var reservas model.Reservas
	err := Db.Where("fecha_in >= ? AND fecha_out <= ? AND user_id = ?", reserva.FechaIn, reserva.FechaOut, reserva.UserId).Find(&reservas).Error
	if err != nil {
		return nil
	}
	return reservas
}

func (c *reservaClient) GetReservasByHotelFechaAndUser(reserva model.Reserva) model.Reservas {
	var reservas model.Reservas
	err := Db.Where("fecha_in >= ? AND fecha_out <= ? AND user_id = ? AND hotel_id = ?", reserva.FechaIn, reserva.FechaOut, reserva.UserId, reserva.HotelId).Find(&reservas).Error
	if err != nil {
		return nil
	}
	return reservas
}
