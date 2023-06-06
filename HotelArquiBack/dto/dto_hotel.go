package dto

type HotelDto struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	CantHabitaciones     int    `json:"CantHabitaciones"`
	CantHabitacionesDisp int    `json:"CantHabitacionesDisp"`
}
type HotelsDto []HotelDto
