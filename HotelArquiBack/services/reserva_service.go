package services

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	reservaClient "HotelArquiSoft/HotelArquiBack/clients/reserva"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
)

type reservaService struct{}

type reservaServiceInterface interface {
	GetReservaById(id int) (dto.ReservaDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
	InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError)
	GetRooms(ReservaDto dto.ReservaDto) bool
	GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError)
	GetReservasByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError)
	GetHotelsByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError)
}

var (
	ReservaService reservaServiceInterface
)

func init() {
	ReservaService = &reservaService{}
}

func (s *reservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {

	var reserva model.Reserva = reservaClient.ReservaClient.GetReservaById(id)
	var reservaDto dto.ReservaDto

	if reserva.ID == 0 {
		return reservaDto, e.NewBadRequestApiError("reserva not found")
	}

	reservaDto.FechaIngreso = reserva.FechaIn
	reservaDto.FechaEgreso = reserva.FechaOut
	reservaDto.HotelId = reserva.HotelId
	reservaDto.UserId = reserva.UserId
	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
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

	reserva = reservaClient.ReservaClient.InsertReserva(reserva)

	reservaDto.Id = reserva.ID

	return reservaDto, nil
}

func (s *reservaService) GetRooms(reservaDto dto.ReservaDto) bool {
	fecha := reservaDto.FechaIngreso

	var reserva model.Reserva

	reserva.FechaIn = reservaDto.FechaIngreso
	reserva.FechaOut = reservaDto.FechaEgreso
	reserva.HotelId = reservaDto.HotelId
	reserva.UserId = reservaDto.UserId

	reservaDto.Id = reserva.ID

	Hotel, _ := HotelService.GetHotelById(reserva.HotelId)

	duracion := reservaDto.FechaEgreso.Sub(fecha)

	dias := int(duracion.Hours() / 24)

	for i := 0; i < dias; i++ {
		if reservaClient.ReservaClient.GetRooms(fecha, reserva) >= Hotel.CantHabitaciones {
			return false
		}
		fecha = fecha.AddDate(0, 0, 1)
	}
	return true
}

func (s *reservaService) GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByUser(userId)
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetReservasByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {

	var reserva model.Reserva

	reserva.FechaIn = reservaDto.FechaIngreso
	reserva.FechaOut = reservaDto.FechaEgreso

	var reservas model.Reservas = reservaClient.ReservaClient.GetReservasByFecha(reserva)
	var reservasDto dto.ReservasDto

	for _, reserva = range reservas {
		var reservaDto dto.ReservaDto

		reservaDto.FechaIngreso = reserva.FechaIn
		reservaDto.FechaEgreso = reserva.FechaOut
		reservaDto.HotelId = reserva.HotelId
		reservaDto.UserId = reserva.UserId
		reservaDto.Id = reserva.ID
		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}

func (s *reservaService) GetHotelsByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {

	hoteles, _ := HotelService.GetHotels()

	var reservasDto dto.ReservasDto

	for _, hotel := range hoteles {
		reservaDto.HotelId = hotel.Id
		if s.GetRooms(reservaDto) {
			reservasDto = append(reservasDto, reservaDto)
		}
	}

	return reservasDto, nil
}
