package controller_test

import (
	e "HotelArquiSoft/HotelArquiBack/Utils"
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	controllerReserva "HotelArquiSoft/HotelArquiBack/controller/reserva"
	"HotelArquiSoft/HotelArquiBack/dto"
	jwtG "HotelArquiSoft/HotelArquiBack/jwt"
	"HotelArquiSoft/HotelArquiBack/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetReservaById(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservaDto := dto.ReservaDto{
		Id:           1,
		UserId:       1,
		HotelId:      1,
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("GetReservaById", 1).Return(mockReservaDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas/:id", controllerReserva.GetReservaById)

	req, err := http.NewRequest("GET", "/reservas/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservaDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservaDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservaById", 1)
}

func TestGetReservas(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasDto := dto.ReservasDto{
		{Id: 1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:           2,
			UserId:       2,
			HotelId:      2,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaService.On("GetReservas").Return(mockReservasDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas", controllerReserva.GetReservas)

	req, err := http.NewRequest(http.MethodGet, "/reservas", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservas")
}

func TestGetReservasByUser(t *testing.T) {
	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasByUserDto := dto.ReservasDto{
		dto.ReservaDto{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaService.On("GetReservasByUser", 1).Return(mockReservasByUserDto, nil)
	services.ReservaService = mockReservaService

	router.GET("/reservas/:user_id", controllerReserva.GetReservasByUser)

	req, err := http.NewRequest(http.MethodGet, "/reservas/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasByUserDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservasByUser", 1)
}

func TestGetReservasByFecha(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasByFechaDto := dto.ReservasDto{
		{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaDto := dto.ReservaDto{
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("GetReservasByFecha", mockReservaDto).Return(mockReservasByFechaDto, nil)
	services.ReservaService = mockReservaService

	router.POST("/reservas/byfecha", controllerReserva.GetReservasByFecha)

	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/byfecha", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
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

	req.Header.Set("Authorization", tokenString)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasByFechaDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservasByFecha", mockReservaDto)
}

func TestGetHotelsByFecha(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservasByFechaDto := dto.ReservasDto{
		dto.ReservaDto{
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
	}
	mockReservaDto := dto.ReservaDto{
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("GetHotelsByFecha", mockReservaDto).Return(mockReservasByFechaDto, nil)
	services.ReservaService = mockReservaService

	router.POST("/reservas/hotelsbyfecha", controllerReserva.GetHotelsByFecha)

	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/hotelsbyfecha", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

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

	req.Header.Set("Authorization", tokenString)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	log.Debug(resp.Code)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasByFechaDto, responseDto)

	mockReservaService.AssertCalled(t, "GetHotelsByFecha", mockReservaDto)
}

func TestReservaInsert(t *testing.T) {

	router := gin.Default()

	mockReservaService := &MockReservaService{}
	mockReservaDto := dto.ReservaDto{

		UserId:       1,
		HotelId:      1,
		FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
		FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
	}
	mockReservaService.On("InsertReserva", mockReservaDto).Return(mockReservaDto, nil)
	services.ReservaService = mockReservaService

	router.POST("/reservas", controllerReserva.ReservaInsert)

	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

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

	req.Header.Set("Authorization", tokenString)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	mockReservaService.AssertCalled(t, "InsertReserva", mockReservaDto)
}

func TestGetReservasByHotelAndFecha(t *testing.T) {

	router := gin.Default()

	FechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	FechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)

	mockReservaService := &MockReservaService{}
	mockReservasByHotelAndFechaDto := dto.ReservasDto{
		{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: FechaI,
			FechaEgreso:  FechaE,
		},
	}
	mockReservaDto := dto.ReservaDto{
		HotelId:      1,
		FechaIngreso: FechaI,
		FechaEgreso:  FechaE,
	}
	mockReservaService.On("GetReservasByHotelAndFecha", mockReservaDto).Return(mockReservasByHotelAndFechaDto, nil)
	services.ReservaService = mockReservaService

	router.POST("/reservas/fechahotel", controllerReserva.GetReservasByHotelAndFecha)

	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/fechahotel", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
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

	req.Header.Set("Authorization", tokenString)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockReservasByHotelAndFechaDto, responseDto)

	mockReservaService.AssertCalled(t, "GetReservasByHotelAndFecha", mockReservaDto)
}

func TestGetReservasByFechaAndUser(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	FechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	FechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	// Set up a mock implementation of the ReservaService
	mockReservaService := &MockReservaService{}
	mockReservasByHotelAndFechaDto := dto.ReservasDto{
		{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: FechaI,
			FechaEgreso:  FechaE,
		},
		// Add more mock data as needed
	}
	mockReservaDto := dto.ReservaDto{
		UserId:       1,
		FechaIngreso: FechaI,
		FechaEgreso:  FechaE,
	}
	mockReservaService.On("GetReservasByFechaAndUser", mockReservaDto).Return(mockReservasByHotelAndFechaDto, nil)
	services.ReservaService = mockReservaService

	// Define the route and handler function
	router.POST("/reservas/fechauser", controllerReserva.GetReservasByFechaAndUser)

	// Create a test request for the route
	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/fechauser", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type header
	req.Header.Set("Content-Type", "application/json")
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

	// Set the Authorization header with the JWT token
	req.Header.Set("Authorization", tokenString)
	// Create a response recorder to capture the response
	resp := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(resp, req)

	// Assert the response status code is http.StatusOK
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body to retrieve the ReservasDto
	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the responseDto matches the expected mockReservasByFechaDto
	assert.Equal(t, mockReservasByHotelAndFechaDto, responseDto)

	// Assert that the ReservaService's GetReservasByFecha method was called with the correct parameter
	mockReservaService.AssertCalled(t, "GetReservasByFechaAndUser", mockReservaDto)
}

func TestGetReservasByHotelAndUser(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	FechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	FechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	// Set up a mock implementation of the ReservaService
	mockReservaService := &MockReservaService{}
	mockReservasByHotelAndFechaDto := dto.ReservasDto{
		{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: FechaI,
			FechaEgreso:  FechaE,
		},
		// Add more mock data as needed
	}
	mockReservaDto := dto.ReservaDto{
		UserId:  1,
		HotelId: 1,
	}
	mockReservaService.On("GetReservasByHotelAndUser", mockReservaDto).Return(mockReservasByHotelAndFechaDto, nil)
	services.ReservaService = mockReservaService

	// Define the route and handler function
	router.POST("/reservas/hoteluser", controllerReserva.GetReservasByHotelAndUser)

	// Create a test request for the route
	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/hoteluser", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type header
	req.Header.Set("Content-Type", "application/json")
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

	// Set the Authorization header with the JWT token
	req.Header.Set("Authorization", tokenString)
	// Create a response recorder to capture the response
	resp := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(resp, req)

	// Assert the response status code is http.StatusOK
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body to retrieve the ReservasDto
	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the responseDto matches the expected mockReservasByFechaDto
	assert.Equal(t, mockReservasByHotelAndFechaDto, responseDto)

	// Assert that the ReservaService's GetReservasByFecha method was called with the correct parameter
	mockReservaService.AssertCalled(t, "GetReservasByHotelAndUser", mockReservaDto)
}

func TestGetReservasByHotelFechaAndUser(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	FechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	FechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	// Set up a mock implementation of the ReservaService
	mockReservaService := &MockReservaService{}
	mockReservasByHotelAndFechaDto := dto.ReservasDto{
		{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: FechaI,
			FechaEgreso:  FechaE,
		},
		// Add more mock data as needed
	}
	mockReservaDto := dto.ReservaDto{
		UserId:       1,
		HotelId:      1,
		FechaIngreso: FechaI,
		FechaEgreso:  FechaE,
	}
	mockReservaService.On("GetReservasByHotelFechaAndUser", mockReservaDto).Return(mockReservasByHotelAndFechaDto, nil)
	services.ReservaService = mockReservaService

	// Define the route and handler function
	router.POST("/reservas/hotelfechauser", controllerReserva.GetReservasByHotelFechaAndUser)

	// Create a test request for the route
	requestBody, err := json.Marshal(mockReservaDto)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "/reservas/hotelfechauser", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type header
	req.Header.Set("Content-Type", "application/json")
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

	// Set the Authorization header with the JWT token
	req.Header.Set("Authorization", tokenString)
	// Create a response recorder to capture the response
	resp := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(resp, req)

	// Assert the response status code is http.StatusOK
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body to retrieve the ReservasDto
	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the responseDto matches the expected mockReservasByFechaDto
	assert.Equal(t, mockReservasByHotelAndFechaDto, responseDto)

	// Assert that the ReservaService's GetReservasByFecha method was called with the correct parameter
	mockReservaService.AssertCalled(t, "GetReservasByHotelFechaAndUser", mockReservaDto)
}

func TestGetReservasByHotel(t *testing.T) {
	router := gin.Default()

	// Set up a mock implementation of the ReservaService
	mockReservaService := &MockReservaService{}
	mockReservasByHotelDto := dto.ReservasDto{
		dto.ReservaDto{
			Id:           1,
			UserId:       1,
			HotelId:      1,
			FechaIngreso: time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC),
			FechaEgreso:  time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC),
		},
		// Add more mock data as needed
	}
	mockReservaService.On("GetReservasByHotel", 1).Return(mockReservasByHotelDto, nil)
	services.ReservaService = mockReservaService

	// Define the route and handler function
	router.GET("/reservas/hotel/:hotel_id", controllerReserva.GetReservasByHotel)

	// Create a test request for the route
	req, err := http.NewRequest(http.MethodGet, "/reservas/hotel/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	resp := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(resp, req)

	// Assert the response status code is http.StatusOK
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body to retrieve the ReservasDto
	var responseDto dto.ReservasDto
	err = json.Unmarshal(resp.Body.Bytes(), &responseDto)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the responseDto matches the expected mockReservasByUserDto
	assert.Equal(t, mockReservasByHotelDto, responseDto)

	// Assert that the ReservaService's GetReservasByUser method was called with the correct parameter
	mockReservaService.AssertCalled(t, "GetReservasByHotel", 1)
}

// Create a mock struct that implements the reservaServiceInterface
type MockReservaService struct {
	mock.Mock
}

func (m *MockReservaService) GetReservaById(id int) (dto.ReservaDto, e.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.ReservaDto), nil
}

func (m *MockReservaService) GetReservas() (dto.ReservasDto, e.ApiError) {
	args := m.Called()
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) InsertReserva(reservaDto dto.ReservaDto) (dto.ReservaDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservaDto), nil
}

func (m *MockReservaService) GetRooms(reservaDto dto.ReservaDto) bool {
	args := m.Called(reservaDto)
	return args.Bool(0)
}

func (m *MockReservaService) GetReservasByUser(userId int) (dto.ReservasDto, e.ApiError) {
	args := m.Called(userId)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetHotelsByFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByHotel(hotelId int) (dto.ReservasDto, e.ApiError) {
	args := m.Called(hotelId)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByHotelAndFecha(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByHotelAndUser(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByFechaAndUser(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}

func (m *MockReservaService) GetReservasByHotelFechaAndUser(reservaDto dto.ReservaDto) (dto.ReservasDto, e.ApiError) {
	args := m.Called(reservaDto)
	return args.Get(0).(dto.ReservasDto), nil
}


