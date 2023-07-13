package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	imagenClient "HotelArquiSoft/HotelArquiBack/clients/imagenes"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type imagenService struct{}

type imagenServiceInterface interface {
	InsertImagen(imageDTO dto.ImageDTO) (dto.ImageDTO, e.ApiError)
	GetImagenesByHotel(hotelId int) (dto.ImagenesDTO, e.ApiError)
}

var (
	ImagenService imagenServiceInterface
)

func init() {
	ImagenService = &imagenService{}
}

func (s *imagenService) InsertImagen(imageDTO dto.ImageDTO) (dto.ImageDTO, e.ApiError) {

	var imagen model.Imagen

	imagen.Nombre = imageDTO.Nombre
	imagen.Contenido = imageDTO.Contenido
	imagen.HotelId = imageDTO.HotelId
	imagen = imagenClient.InsertImagen(imagen)
	return imageDTO, nil
}

func (s *imagenService) GetImagenesByHotel(hotelId int) (dto.ImagenesDTO, e.ApiError) {

	var imagenes model.Imagenes = imagenClient.GetImagenesByHotel(hotelId)
	var imagenesDto dto.ImagenesDTO

	for _, imagen := range imagenes {
		var imagenDto dto.ImageDTO
		imagenDto.HotelId = imagen.HotelId
		imagenDto.Contenido = imagen.Contenido
		imagenDto.Nombre = imagen.Nombre
		imagenesDto = append(imagenesDto, imagenDto)
	}

	return imagenesDto, nil
}
