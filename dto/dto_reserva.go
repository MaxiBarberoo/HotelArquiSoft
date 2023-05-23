package dto

type ReservaDto struct {
	Id           int `json:"id"`
	UserId       int `json:"user_id,omitempty"`
	HotelId      int `json:"hotel_id,omitempty"`
	FechaIngreso int `json:"fecha_ingreso"`
	FechaEgreso  int `json:"fecha_egreso"`
}

type ReservasDto []ReservaDto
