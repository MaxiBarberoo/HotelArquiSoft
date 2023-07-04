package dto

type AmenitieHotelDto struct {
	Id         int    `json:"id"`
	Tipo       string `json:"tipo"`
	AmenitieId int    `json:"user_id,omitempty"`
	HotelId    int    `json:"hotel_id,omitempty"`
}

type AmenitiesHotels []AmenitieHotelDto
