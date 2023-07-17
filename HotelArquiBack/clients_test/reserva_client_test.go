package clients_test

import (
	clients "HotelArquiSoft/HotelArquiBack/clients/reserva"
	"HotelArquiSoft/HotelArquiBack/model"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import SQLite dialect for GORM
	"github.com/stretchr/testify/assert"
)

func TestReservaClient(t *testing.T) {
	// Inicializar la base de datos en memoria para las pruebas
	Db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Migrar el modelo de Reserva para crear la tabla en la base de datos en memoria
	Db.AutoMigrate(&model.Reserva{})
	Db.AutoMigrate(&model.Hotel{})

	// Asignar la base de datos en memoria al cliente de Reservas
	clients.Db = Db

	hotel := model.Hotel{
		Nombre:      "Luxury",
		CantHab:     10,
		Descripcion: "Hotel Mock",
	}

	Db.Create(&hotel)
	// Insertar una Reserva de ejemplo para probar las funciones
	fechaI := time.Date(2023, time.July, 15, 0, 0, 0, 0, time.UTC)
	fechaE := time.Date(2023, time.July, 17, 0, 0, 0, 0, time.UTC)
	reserva := model.Reserva{
		FechaIn:  fechaI, // Fecha de salida dentro de 3 días
		FechaOut: fechaE,
		UserId:   1, // ID de usuario de ejemplo
		HotelId:  1, // ID de hotel de ejemplo
	}
	insertedReserva := clients.ReservaClient.InsertReserva(reserva)

	// Probar la función GetReservaById
	foundReserva := clients.ReservaClient.GetReservaById(insertedReserva.ID)
	assert.Equal(t, insertedReserva.ID, foundReserva.ID, "Las Reservas deben coincidir")

	// Probar la función GetReservas
	reservas := clients.ReservaClient.GetReservas()
	assert.NotEmpty(t, reservas, "Debería haber al menos una Reserva en la base de datos")
	assert.Equal(t, reserva.UserId, reservas[0].UserId)
	assert.Equal(t, reserva.HotelId, reservas[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservas[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservas[0].FechaOut)
	// Probar la función GetReservasByUser
	reservasByUser := clients.ReservaClient.GetReservasByUser(reserva.UserId)
	assert.NotEmpty(t, reservasByUser, "Debería haber al menos una Reserva para el usuario especificado")
	assert.Equal(t, reserva.UserId, reservasByUser[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByUser[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByUser[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByUser[0].FechaOut)

	// Probar la función GetReservasByFecha
	reservasByFecha := clients.ReservaClient.GetReservasByFecha(reserva)
	assert.NotEmpty(t, reservasByFecha, "Debería haber al menos una Reserva para la fecha especificada")
	assert.Equal(t, reserva.UserId, reservasByFecha[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByFecha[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByFecha[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByFecha[0].FechaOut)
	// Probar la función GetRooms
	countRooms := clients.ReservaClient.GetRooms(time.Now(), reserva)
	assert.Equal(t, 1, countRooms, "Debería haber una habitación reservada para la fecha y el hotel especificados")

	reservasByHotelAndFecha := clients.ReservaClient.GetReservasByHotelAndFecha(reserva)
	assert.NotEmpty(t, reservasByHotelAndFecha, "Debería haber al menos una Reserva para la fecha y hotel especificada")
	assert.Equal(t, reserva.UserId, reservasByHotelAndFecha[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByHotelAndFecha[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByHotelAndFecha[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByHotelAndFecha[0].FechaOut)

	reservasByHotelAndUser := clients.ReservaClient.GetReservasByHotelAndUser(reserva)
	assert.NotEmpty(t, reservasByHotelAndUser, "Debería haber al menos una Reserva para la user y hotel especificada")
	assert.Equal(t, reserva.UserId, reservasByHotelAndUser[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByHotelAndUser[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByHotelAndUser[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByHotelAndUser[0].FechaOut)

	reservasByFechaAndUser := clients.ReservaClient.GetReservasByFechaAndUser(reserva)
	assert.NotEmpty(t, reservasByFechaAndUser, "Debería haber al menos una Reserva para la user y fecha especificada")
	assert.Equal(t, reserva.UserId, reservasByFechaAndUser[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByFechaAndUser[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByFechaAndUser[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByFechaAndUser[0].FechaOut)
	reservasByHotelFechaAndUser := clients.ReservaClient.GetReservasByHotelFechaAndUser(reserva)
	assert.NotEmpty(t, reservasByHotelFechaAndUser, "Debería haber al menos una Reserva para la hotel, user y fecha especificada")
	assert.Equal(t, reserva.UserId, reservasByHotelFechaAndUser[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByHotelFechaAndUser[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByHotelFechaAndUser[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByHotelFechaAndUser[0].FechaOut)

	reservasByHotel := clients.ReservaClient.GetReservasByHotel(reserva.HotelId)
	assert.NotEmpty(t, reservasByHotel, "Debería haber al menos una Reserva para el hotel especificado")
	assert.Equal(t, reserva.UserId, reservasByHotel[0].UserId)
	assert.Equal(t, reserva.HotelId, reservasByHotel[0].HotelId)
	assert.Equal(t, reserva.FechaIn, reservasByHotel[0].FechaIn)
	assert.Equal(t, reserva.FechaOut, reservasByHotel[0].FechaOut)
	err = Db.Close()
	if err != nil {
		t.Fatalf("Error al cerrar la base de datos: %v", err)
	}
}
