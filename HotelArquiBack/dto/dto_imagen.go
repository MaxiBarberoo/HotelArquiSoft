package dto

type ImageDTO struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Contenido []byte `json:"contenido"`
	HotelId   int    `json:"hotel_id"`
}

type ImagenesDTO []ImageDTO
