package end_to_end_tests

import (
	reservaClient "HotelArquiSoft/HotelArquiBack/clients/reserva"
	reservaController "HotelArquiSoft/HotelArquiBack/controller/reserva"
	"HotelArquiSoft/HotelArquiBack/dto"
	jwtG "HotelArquiSoft/HotelArquiBack/jwt"
	"HotelArquiSoft/HotelArquiBack/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Use the SQLite driver for testing
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetReservasByFecha(t *testing.T) {
	// Create a new SQLite in-memory database for testing
	Db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	fechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	fechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	var mockUserDto dto.UserDto
	mockUserDto.Id = 1
	mockUserDto.UserEmail = "mock@mock.com"
	mockUserDto.Tipo = 0
	mockUserDto.FirstName = "Mock"
	mockUserDto.LastName = "Mock"

	tokenString, err := jwtG.GenerateUserToken(mockUserDto)

	if err != nil {
		t.Fatal(err)
	}
	// Auto-migrate the database to create the Reserva table
	Db.AutoMigrate(&model.Reserva{})

	// Replacing the client reservaClient with the in-memory SQLite database
	reservaClient.Db = Db

	// Create a new Gin router
	router := gin.Default()

	// Configure the route for the GetReservasByFecha function
	router.POST("/reservas/byfecha", reservaController.GetReservasByFecha)

	// Create a test reservation to send in the request
	reserva := dto.ReservaDto{
		FechaIngreso: fechaI,
		FechaEgreso:  fechaE,
	}

	// Convert the reservation to JSON to send in the request body
	reservaJSON, err := json.Marshal(reserva)
	if err != nil {
		t.Fatalf("Error converting the reservation to JSON: %v", err)
	}

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/reservas/byfecha", bytes.NewBuffer(reservaJSON))
	if err != nil {
		t.Fatalf("Error creating the HTTP request: %v", err)
	}
	req.Header.Set("Authorization", tokenString)
	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a test HTTP response recorder
	rec := httptest.NewRecorder()

	// Configure the in-memory SQLite database to return test data
	mockReservas := []model.Reserva{
		{
			ID:       1,
			FechaIn:  fechaI,
			FechaOut: fechaE,
			UserId:   1,
			HotelId:  1,
		},
		{
			ID:       2,
			FechaIn:  fechaI,
			FechaOut: fechaE,
			UserId:   2,
			HotelId:  2,
		},
	}

	// Insert the mock reservations into the in-memory SQLite database
	for _, r := range mockReservas {
		Db.Create(&r)
	}

	// Execute the request on the router
	router.ServeHTTP(rec, req)

	// Verify that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify that the response has a valid JSON format
	var response dto.ReservasDto
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify that the response contains the expected reservations based on the provided date range
	assert.Len(t, response, 2)
	assert.Equal(t, mockReservas[0].ID, response[0].Id)
	assert.Equal(t, mockReservas[0].FechaIn, response[0].FechaIngreso)
	assert.Equal(t, mockReservas[0].FechaOut, response[0].FechaEgreso)
	assert.Equal(t, mockReservas[0].UserId, response[0].UserId)
	assert.Equal(t, mockReservas[0].HotelId, response[0].HotelId)

	assert.Equal(t, mockReservas[1].ID, response[1].Id)
	assert.Equal(t, mockReservas[1].FechaIn, response[1].FechaIngreso)
	assert.Equal(t, mockReservas[1].FechaOut, response[1].FechaEgreso)
	assert.Equal(t, mockReservas[1].UserId, response[1].UserId)
	assert.Equal(t, mockReservas[1].HotelId, response[1].HotelId)

	err = Db.Close()
	if err != nil {
		t.Fatalf("Error al cerrar la base de datos: %v", err)
	}
}
