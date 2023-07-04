package dto

type AmenitieDto struct {
	Id   int    `json:"id"`
	Tipo string `json:"tipo"`
}

type AmenitiesDto []AmenitieDto
