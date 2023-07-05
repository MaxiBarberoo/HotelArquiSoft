package dto

type AmenitieHotelDto struct {
	Id         int `json:"id"`
	AmenitieId int `json:"amenitie_id"`
	HotelId    int `json:"hotel_id"`
}

type AmenitiesHotels []AmenitieHotelDto
