package services

import (
	e "HotelArquiSoft/Utils"
	reservaClient "HotelArquiSoft/clients/reserva"
	"HotelArquiSoft/dto"
	"HotelArquiSoft/model"
)

type reservaService struct{}

type reservaServiceInterface interface {
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetRooms(ReservaDto dto.ReservaDto, Hotel model.Hotel) bool
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("reserva not found")
	}

	reservaDto.FechaIngreso = reserva.FechaIn
	reservaDto.FechaEgreso = reserva.FechaOut
	reservaDto.HotelId = reserva.HotelId
	reservaDto.UserId = reserva.UserId

	return reservaDto, nil
}

func (s *reservaService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId

		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva

	reserva.FechaIn = reservaDto.FechaIngreso
	reserva.FechaOut = reservaDto.FechaEgreso
	reserva.HotelId = reservaDto.HotelId
	reserva.UserId = reservaDto.UserId

	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetRooms(reservaDto dto.ReservaDto, Hotel model.Hotel) bool {
	fecha := reservaDto.FechaIngreso

	var reserva model.Reserva

	reserva.FechaIn = reservaDto.FechaIngreso
	reserva.FechaOut = reservaDto.FechaEgreso
	reserva.HotelId = reservaDto.HotelId
	reserva.UserId = reservaDto.UserId

	reservaDto.Id = reserva.ID

	duracion := reservaDto.FechaEgreso.Sub(fecha)

	dias := int(duracion.Hours() / 24)

	for i := 0; i < dias; i++ {
		if reservaClient.GetRooms(fecha, reserva) > Hotel.CantHab {
			return false
		}
		fecha = fecha.AddDate(0, 0, 1)
	}
	return true
}
