package model

import "time"

type Reserva struct {
	ID       int
	FechaIn  time.Time `gorm:"type:DATE"`
	FechaOut time.Time `gorm:"type:DATE"`
	UserId   int       `gorm:"foreignKey"`
	HotelId  int       `gorm:"foreignKey"`
}

type Reservas []Reserva
