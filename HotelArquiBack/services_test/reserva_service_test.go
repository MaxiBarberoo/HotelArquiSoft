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

// Mock for the ReservaClientInterface
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

// Test GetReservaById
func TestGetReservaById(t *testing.T) {
	// Arrange
	mockClient := new(mockReservaClient)

	expectedReserva := model.Reserva{
		ID:       1,
		FechaIn:  time.Now(),
		FechaOut: time.Now().Add(time.Hour * 24),
		UserId:   123,
		HotelId:  456,
	}

	mockClient.On("GetReservaById", 1).Return(expectedReserva)

	// Set the mocked client for the ReservaService
	clients.ReservaClient = mockClient

	// Act
	reservaDto, err := services.ReservaService.GetReservaById(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedReserva.ID, reservaDto.Id)
	assert.Equal(t, expectedReserva.FechaIn, reservaDto.FechaIngreso)
	assert.Equal(t, expectedReserva.FechaOut, reservaDto.FechaEgreso)
	assert.Equal(t, expectedReserva.UserId, reservaDto.UserId)
	assert.Equal(t, expectedReserva.HotelId, reservaDto.HotelId)
	mockClient.AssertExpectations(t)
}

func TestGetReservas(t *testing.T) {
	// Arrange
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

	// Set the mocked client for the ReservaService
	clients.ReservaClient = mockClient

	// Act
	reservasDto, err := services.ReservaService.GetReservas()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	// Check the first reserva
	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	// Check the second reserva
	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)

}

func TestGetRooms(t *testing.T) {
	// Arrange
	mockClient := new(mockReservaClient)
	mockHotel := new(mockHotelService)
	// Sample data for the reservaDto
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

	// Sample data for the reserva (which will be passed to the client method)
	reserva := model.Reserva{
		FechaIn:  fechaIngreso,
		FechaOut: fechaEgreso,
		UserId:   reservaDto.UserId,
		HotelId:  reservaDto.HotelId,
		ID:       0,
	}

	// Set the expected count returned by the client method
	mockClient.On("GetRooms", fechaIngreso, reserva).Return(5)
	mockHotel.On("GetHotelById", 456).Return(expectedHotel, nil)
	// Create a new instance of the service with the mock client
	services.HotelService = mockHotel
	clients.ReservaClient = mockClient

	// Act
	result := services.ReservaService.GetRooms(reservaDto)

	// Assert
	assert.True(t, result)

	mockClient.AssertExpectations(t)
	mockHotel.AssertExpectations(t)
}

func TestGetReservasByUser(t *testing.T) {
	// Arrange
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

	// Create a new instance of the service with the mock client
	clients.ReservaClient = mockClient

	// Act
	reservasDto, err := services.ReservaService.GetReservasByUser(userId)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, reservasDto, 2)

	// Check the first reserva
	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	// Check the second reserva
	assert.Equal(t, expectedReservas[1].ID, reservasDto[1].Id)
	assert.Equal(t, expectedReservas[1].FechaIn, reservasDto[1].FechaIngreso)
	assert.Equal(t, expectedReservas[1].FechaOut, reservasDto[1].FechaEgreso)
	assert.Equal(t, expectedReservas[1].UserId, reservasDto[1].UserId)
	assert.Equal(t, expectedReservas[1].HotelId, reservasDto[1].HotelId)

	mockClient.AssertExpectations(t)
}

func TestGetReservasByFecha(t *testing.T) {
	// Arrange
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

	// Create a new instance of the service with the mock client
	clients.ReservaClient = mockClient

	// Act
	reservasDto, err := services.ReservaService.GetReservasByFecha(newReservaDto)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, reservasDto, 1)

	// Check the first reserva
	assert.Equal(t, expectedReservas[0].ID, reservasDto[0].Id)
	assert.Equal(t, expectedReservas[0].FechaIn, reservasDto[0].FechaIngreso)
	assert.Equal(t, expectedReservas[0].FechaOut, reservasDto[0].FechaEgreso)
	assert.Equal(t, expectedReservas[0].UserId, reservasDto[0].UserId)
	assert.Equal(t, expectedReservas[0].HotelId, reservasDto[0].HotelId)

	mockClient.AssertExpectations(t)
}

/*
func TestGetHotelsByFecha(t *testing.T) {
	// Arrange
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

	// Mock the GetRooms function

	// Mock the GetHotelById function
	mockClient.On("GetRooms", fechaIngreso, reserva).Return(5)
	mockHotel.On("GetHotelById", 1).Return(expectedHotel1, nil)
	mockHotel.On("GetHotels").Return(expectedHotels, nil)

	// Set the mocked hotel and reserva clients for the services package
	services.HotelService = mockHotel
	clients.ReservaClient = mockClient

	// Act
	hotelsDto, err := services.ReservaService.GetHotelsByFecha(reservaDto)
	// Assert
	assert.NoError(t, err)
	assert.Len(t, hotelsDto, 1)
	assert.Equal(t, expectedHotel1.Id, hotelsDto[0].Id)

	// Verify that the GetRooms function is called twice with the correct parameters
	mockClient.AssertExpectations(t)

	// Verify that the GetHotelById function is called for both hotels
	mockHotel.AssertExpectations(t)
}*/
