package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"HotelArquiSoft/HotelArquiBack/dto"
	"HotelArquiSoft/HotelArquiBack/model"
	"HotelArquiSoft/HotelArquiBack/services"
)

// TestGetReservaById tests the GetReservaById function of ReservaService.
func TestGetReservaById(t *testing.T) {
	// Mock the GetReservaById function of the reservaClient
	mockGetReservaById(1)

	// Set the mock function as the implementation of GetReservaById in the client

	// Call the GetReservaById function with a specific ID
	reservaDto, err := services.ReservaService.GetReservaById(1)

	// Assert that no error occurred
	assert.Nil(t, err)

	// Assert the expected values of the reservation DTO
	expectedReservaDto := dto.ReservaDto{
		Id:           1,
		UserId:       1,
		HotelId:      1,
		FechaIngreso: reservaDto.FechaIngreso,
		FechaEgreso:  reservaDto.FechaEgreso,
	}
	assert.Equal(t, expectedReservaDto, reservaDto)
}

func mockGetReservaById(id int) model.Reserva {
	// Return a mock reservation based on the provided ID
	return model.Reserva{
		ID:       id,
		FechaIn:  time.Now(),
		FechaOut: time.Now().AddDate(0, 0, 1),
		UserId:   1,
		HotelId:  1,
	}
}
