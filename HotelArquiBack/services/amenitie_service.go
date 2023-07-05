package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	amenitieClient "HotelArquiSoft/HotelArquiBack/clients/amenitie"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type amenitieService struct{}

type amenitieServiceInterface interface {
	GetAmenitieById(id int) (dto.AmenitieDto, e.ApiError)
	GetAmenities() (dto.AmenitiesDto, e.ApiError)
}

var (
	AmenitieService amenitieServiceInterface
)

func init() {
	AmenitieService = &amenitieService{}
}

func (s *amenitieService) GetAmenitieById(id int) (dto.AmenitieDto, e.ApiError) {

	var amenitie model.Amenitie = amenitieClient.GetAmenitieById(id)
	var amenitieDto dto.AmenitieDto

	if amenitie.ID == 0 {
		return amenitieDto, e.NewBadRequestApiError("user not found")
	}

	amenitieDto.Tipo = amenitie.Tipo
	amenitieDto.Id = amenitie.ID

	return amenitieDto, nil
}

func (s *amenitieService) GetAmenities() (dto.AmenitiesDto, e.ApiError) {

	var amenities model.Amenities = amenitieClient.GetAmenities()
	var amenitiesDto dto.AmenitiesDto

	for _, amenitie := range amenities {
		var amenitieDto dto.AmenitieDto
		amenitieDto.Tipo = amenitie.Tipo
		amenitieDto.Id = amenitie.ID
		amenitiesDto = append(amenitiesDto, amenitieDto)
	}

	return amenitiesDto, nil
}
