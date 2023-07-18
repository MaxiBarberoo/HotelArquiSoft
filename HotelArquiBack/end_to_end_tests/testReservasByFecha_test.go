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
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetReservasByFecha(t *testing.T) {
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
	Db.AutoMigrate(&model.Reserva{})

	reservaClient.Db = Db

	router := gin.Default()

	router.POST("/reservas/byfecha", reservaController.GetReservasByFecha)

	reserva := dto.ReservaDto{
		FechaIngreso: fechaI,
		FechaEgreso:  fechaE,
	}

	reservaJSON, err := json.Marshal(reserva)
	if err != nil {
		t.Fatalf("Error converting the reservation to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/reservas/byfecha", bytes.NewBuffer(reservaJSON))
	if err != nil {
		t.Fatalf("Error creating the HTTP request: %v", err)
	}
	req.Header.Set("Authorization", tokenString)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

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

	for _, r := range mockReservas {
		Db.Create(&r)
	}

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response dto.ReservasDto
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

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
