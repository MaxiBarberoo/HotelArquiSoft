package dto

type HotelDto struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	CantHabitaciones int    `json:"cantHabitaciones"`
}
type HotelsDto []HotelDto
