package dto

import "time"

type ReservaDto struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id,omitempty"`
	HotelId      int       `json:"hotel_id,omitempty"`
	FechaIngreso time.Time `json:"fecha_ingreso"`
	FechaEgreso  time.Time `json:"fecha_egreso"`
}

type ReservasDto []ReservaDto
