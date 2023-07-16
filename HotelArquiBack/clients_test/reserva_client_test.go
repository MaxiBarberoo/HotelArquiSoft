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
	reserva := model.Reserva{
		FechaIn:  time.Now(),
		FechaOut: time.Now().AddDate(0, 0, 3), // Fecha de salida dentro de 3 días
		UserId:   1,                           // ID de usuario de ejemplo
		HotelId:  1,                           // ID de hotel de ejemplo
	}
	insertedReserva := clients.ReservaClient.InsertReserva(reserva)

	// Probar la función GetReservaById
	foundReserva := clients.ReservaClient.GetReservaById(insertedReserva.ID)
	assert.Equal(t, insertedReserva.ID, foundReserva.ID, "Las Reservas deben coincidir")

	// Probar la función GetReservas
	reservas := clients.ReservaClient.GetReservas()
	assert.NotEmpty(t, reservas, "Debería haber al menos una Reserva en la base de datos")

	// Probar la función GetReservasByUser
	reservasByUser := clients.ReservaClient.GetReservasByUser(reserva.UserId)
	assert.NotEmpty(t, reservasByUser, "Debería haber al menos una Reserva para el usuario especificado")

	// Probar la función GetReservasByFecha
	reservasByFecha := clients.ReservaClient.GetReservasByFecha(reserva)
	assert.NotEmpty(t, reservasByFecha, "Debería haber al menos una Reserva para la fecha especificada")

	// Probar la función GetRooms
	countRooms := clients.ReservaClient.GetRooms(time.Now(), reserva)
	assert.Equal(t, 1, countRooms, "Debería haber una habitación reservada para la fecha y el hotel especificados")

	err = Db.Close()
	if err != nil {
		t.Fatalf("Error al cerrar la base de datos: %v", err)
	}
}
