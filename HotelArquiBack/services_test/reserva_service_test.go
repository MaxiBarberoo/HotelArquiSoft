package services_test

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	clients "HotelArquiSoft/HotelArquiBack/clients/reserva"
	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
	"HotelArquiSoft/HotelArquiBack/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockReservaClient struct {
	mock.Mock
}

func (m *mockReservaClient) GetReservaById(id int) model.Reserva {
	args := m.Called(id)
	return args.Get(0).(model.Reserva)
}

func (m *mockReservaClient) GetReservas() model.Reservas {
	args := m.Called()
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) InsertReserva(reserva model.Reserva) model.Reserva {
	args := m.Called(reserva)
	return args.Get(0).(model.Reserva)
}

func (m *mockReservaClient) GetRooms(fecha time.Time, reserva model.Reserva) int {
	args := m.Called(fecha, reserva)
	return args.Int(0)
}

func (m *mockReservaClient) GetReservasByUser(userId int) model.Reservas {
	args := m.Called(userId)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByFecha(reserva model.Reserva) model.Reservas {
	args := m.Called(reserva)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByHotelAndFecha(reserva model.Reserva) model.Reservas {
	args := m.Called(reserva)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByHotelAndUser(reserva model.Reserva) model.Reservas {
	args := m.Called(reserva)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByFechaAndUser(reserva model.Reserva) model.Reservas {
	args := m.Called(reserva)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByHotelFechaAndUser(reserva model.Reserva) model.Reservas {
	args := m.Called(reserva)
	return args.Get(0).(model.Reservas)
}

func (m *mockReservaClient) GetReservasByHotel(hotelId int) model.Reservas {
	args := m.Called(hotelId)
	return args.Get(0).(model.Reservas)
}

type mockHotelService struct {
	mock.Mock
}

func (m *mockHotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.HotelDto), nil
}

func (m *mockHotelService) GetHotels() (dto.HotelsDto, e.ApiError) {
	args := m.Called()
	return args.Get(0).(dto.HotelsDto), nil
}

func (m *mockHotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	args := m.Called(hotelDto)
	return args.Get(0).(dto.HotelDto), nil
}

func TestGetReservaById(t *testing.T) {
	mockClient := new(mockReservaClient)

	expectedReserva := model.Reserva{
		ID:       1,
		FechaIn:  time.Now(),
		FechaOut: time.Now().Add(time.Hour * 24),
		UserId:   123,
		HotelId:  456,
	}

	mockClient.On("GetReservaById", 1).Return(expectedReserva)

	clients.ReservaClient = mockClient

	reservaDto, err := services.ReservaService.GetReservaById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedReserva.ID, reservaDto.Id)
	assert.Equal(t, expectedReserva.FechaIn, reservaDto.FechaIngreso)
	assert.Equal(t, expectedReserva.FechaOut, reservaDto.FechaEgreso)
	assert.Equal(t, expectedReserva.UserId, reservaDto.UserId)
	assert.Equal(t, expectedReserva.HotelId, reservaDto.HotelId)
	mockClient.AssertExpectations(t)
}

func TestGetReservas(t *testing.T) {
	mockClient := new(mockReservaClient)

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  time.Now(),
			FechaOut: time.Now().Add(time.Hour * 24),
			UserId:   123,
			HotelId:  456,
		},
		{
			ID:       2,
			FechaIn:  time.Now().Add(time.Hour * 24),
			FechaOut: time.Now().Add(time.Hour * 48),
			UserId:   789,
			HotelId:  123,
		},
	}

	mockClient.On("GetReservas").Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservas()

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)

}

func TestGetRooms(t *testing.T) {
	mockClient := new(mockReservaClient)
	mockHotel := new(mockHotelService)
	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	reservaDto := dto.ReservaDto{
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
		UserId:       123,
		HotelId:      456,
	}

	expectedHotel := dto.HotelDto{
		Id:               456,
		Name:             "Hotel Test",
		CantHabitaciones: 10,
		Desc:             "Test description",
	}

	reserva := model.Reserva{
		FechaIn:  fechaIngreso,
		FechaOut: fechaEgreso,
		UserId:   reservaDto.UserId,
		HotelId:  reservaDto.HotelId,
		ID:       0,
	}

	mockClient.On("GetRooms", fechaIngreso, reserva).Return(5)
	mockHotel.On("GetHotelById", 456).Return(expectedHotel, nil)
	services.HotelService = mockHotel
	clients.ReservaClient = mockClient

	result := services.ReservaService.GetRooms(reservaDto)

	assert.True(t, result)

	mockClient.AssertExpectations(t)
	mockHotel.AssertExpectations(t)
}

func TestGetReservasByUser(t *testing.T) {
	mockClient := new(mockReservaClient)
	userId := 123

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  time.Now(),
			FechaOut: time.Now().Add(time.Hour * 24),
			UserId:   123,
			HotelId:  456,
		},
		{
			ID:       2,
			FechaIn:  time.Now().Add(time.Hour * 24),
			FechaOut: time.Now().Add(time.Hour * 48),
			UserId:   123,
			HotelId:  789,
		},
	}

	mockClient.On("GetReservasByUser", userId).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByUser(userId)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByFecha(t *testing.T) {
	mockClient := new(mockReservaClient)

	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	newReservaDto := dto.ReservaDto{
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
	}

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  fechaIngreso,
			FechaOut: fechaEgreso,
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByFecha", mock.AnythingOfType("model.Reserva")).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByFecha(newReservaDto)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetHotelsByFecha(t *testing.T) {
	mockHotel := new(mockHotelService)
	mockClient := new(mockReservaClient)
	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	expectedHotel1 := dto.HotelDto{
		Id:               1,
		Name:             "Hotel 1",
		CantHabitaciones: 10,
		Desc:             "Description 1",
	}

	expectedHotels := dto.HotelsDto{
		{
			Id:               1,
			Name:             "Hotel 1",
			CantHabitaciones: 10,
			Desc:             "Description 1",
		},
	}
	reservaDto := dto.ReservaDto{
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
		HotelId:      expectedHotel1.Id,
	}

	reserva := model.Reserva{
		FechaIn:  fechaIngreso,
		FechaOut: fechaEgreso,
		HotelId:  expectedHotel1.Id,
	}

	mockClient.On("GetRooms", fechaIngreso, reserva).Return(5)
	mockHotel.On("GetHotelById", 1).Return(expectedHotel1, nil)
	mockHotel.On("GetHotels").Return(expectedHotels, nil)

	services.HotelService = mockHotel
	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetHotelsByFecha(reservaDto)
	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)
	assert.Equal(t, expectedHotel1.Id, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)

	mockHotel.AssertExpectations(t)
}

func TestGetReservasByHotelAndFecha(t *testing.T) {
	mockClient := new(mockReservaClient)

	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	newReservaDto := dto.ReservaDto{
		HotelId:      456,
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
	}

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  fechaIngreso,
			FechaOut: fechaEgreso,
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByHotelAndFecha", mock.AnythingOfType("model.Reserva")).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByHotelAndFecha(newReservaDto)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByHotelAndUser(t *testing.T) {
	mockClient := new(mockReservaClient)

	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	newReservaDto := dto.ReservaDto{
		HotelId: 456,
		UserId:  123,
	}

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  fechaIngreso,
			FechaOut: fechaEgreso,
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByHotelAndUser", mock.AnythingOfType("model.Reserva")).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByHotelAndUser(newReservaDto)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByFechaAndUser(t *testing.T) {
	mockClient := new(mockReservaClient)

	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	newReservaDto := dto.ReservaDto{
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
		UserId:       123,
	}

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  fechaIngreso,
			FechaOut: fechaEgreso,
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByFechaAndUser", mock.AnythingOfType("model.Reserva")).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByFechaAndUser(newReservaDto)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByHotelFechaAndUser(t *testing.T) {
	mockClient := new(mockReservaClient)

	fechaIngreso := time.Now()
	fechaEgreso := fechaIngreso.Add(time.Hour * 24)
	newReservaDto := dto.ReservaDto{
		FechaIngreso: fechaIngreso,
		FechaEgreso:  fechaEgreso,
		UserId:       123,
		HotelId:      456,
	}

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  fechaIngreso,
			FechaOut: fechaEgreso,
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByHotelFechaAndUser", mock.AnythingOfType("model.Reserva")).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByHotelFechaAndUser(newReservaDto)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByHotel(t *testing.T) {
	mockClient := new(mockReservaClient)
	hotelId := 456

	expectedReservas := model.Reservas{
		{
			ID:       1,
			FechaIn:  time.Now(),
			FechaOut: time.Now().Add(time.Hour * 24),
			UserId:   123,
			HotelId:  456,
		},
		{
			ID:       2,
			FechaIn:  time.Now().Add(time.Hour * 24),
			FechaOut: time.Now().Add(time.Hour * 48),
			UserId:   123,
			HotelId:  456,
		},
	}

	mockClient.On("GetReservasByHotel", hotelId).Return(expectedReservas)

	clients.ReservaClient = mockClient

	reservasDto, err := services.ReservaService.GetReservasByHotel(hotelId)

	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)
}
