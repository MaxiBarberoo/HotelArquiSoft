package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	amenitieHotelClient "HotelArquiSoft/HotelArquiBack/clients/amenitiehotel"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type amenitieHotelService struct{}

type amenitieHotelServiceInterface interface {
	AssignAmenitieToHotel(amenitieHotelDto dto.AmenitieHotelDto) (dto.AmenitieHotelDto, e.ApiError)
	SearchAmenitiesByHotel(hotelId int) (dto.AmenitiesHotels, e.ApiError)
}

var (
	AmenitieHotelService amenitieHotelServiceInterface
)

func init() {
	AmenitieHotelService = &amenitieHotelService{}
}

func (s *amenitieHotelService) AssignAmenitieToHotel(amenitieHotelDto dto.AmenitieHotelDto) (dto.AmenitieHotelDto, e.ApiError) {

	var amenitieHotel model.AmenitieHotel

	amenitieHotel.HotelId = amenitieHotelDto.HotelId
	amenitieHotel.AmenitieId = amenitieHotelDto.AmenitieId

	amenitieHotel = amenitieHotelClient.AssignAmenitieToHotel(amenitieHotel)
	amenitieHotelDto.Id = amenitieHotel.ID

	return amenitieHotelDto, nil

}

func (s *amenitieHotelService) SearchAmenitiesByHotel(hotelId int) (dto.AmenitiesHotels, e.ApiError) {
	var amenitiesHotels model.AmenitiesHotels = amenitieHotelClient.SearchAmenitiesByHotel(hotelId)
	var amenitiesHotelsDto dto.AmenitiesHotels

	for _, amenitieHotel := range amenitiesHotels {
		var amenitieHotelDto dto.AmenitieHotelDto
		amenitieHotelDto.AmenitieId = amenitieHotel.AmenitieId
		amenitiesHotelsDto = append(amenitiesHotelsDto, amenitieHotelDto)
	}

	return amenitiesHotelsDto, nil
}
