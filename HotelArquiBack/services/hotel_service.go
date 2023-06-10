package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	hotelClient "HotelArquiSoft/HotelArquiBack/clients/hotel"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("user not found")
	}

	hotelDto.Name = hotel.Nombre
	hotelDto.CantHabitaciones = hotel.CantHab
	hotelDto.Id = hotel.ID

	return hotelDto, nil
}

func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.Name = hotel.Nombre
		hotelDto.CantHabitaciones = hotel.CantHab
		hotelDto.Id = hotel.ID
		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Nombre = hotelDto.Name
	hotel.CantHab = hotelDto.CantHabitaciones
	hotel = hotelClient.InsertHotel(hotel)
	hotelDto.Id = hotel.ID

	return hotelDto, nil
}
