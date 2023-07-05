package db

import (
	amenitieClient "HotelArquiSoft/HotelArquiBack/clients/amenitie"
	amenitieHotelClient "HotelArquiSoft/HotelArquiBack/clients/amenitiehotel"
	hotelClient "HotelArquiSoft/HotelArquiBack/clients/hotel"
	reservaClient "HotelArquiSoft/HotelArquiBack/clients/reserva"
	userClient "HotelArquiSoft/HotelArquiBack/clients/user"
	"HotelArquiSoft/HotelArquiBack/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func insertInitialData() {
	// Insert users
	user := model.User{
		FirstName: "Admin",
		LastName:  "Admin",
		Email:     "admin@admin.com",
		Password:  "password123",
		Tipo:      1,
	}
	if err := db.Create(&user).Error; err != nil {
		log.Error("Failed to insert user:", err.Error())
	}

	// Insert hotels
	hotels := []model.Hotel{
		{
			Nombre:      "Luxury",
			CantHab:     10,
			Descripcion: "Un hotel boutique elegante y moderno, donde el lujo se combina con un servicio impecable y una ubicación privilegiada. Disfruta de habitaciones exquisitamente decoradas, gastronomía de clase mundial y una atención personalizada que hará de tu estancia una experiencia inolvidable.",
		},
		{
			Nombre:      "Grand Hotel",
			CantHab:     10,
			Descripcion: "Sumérgete en la serenidad de nuestro hotel de playa, donde las suaves olas acarician la costa y el sol brilla en un cielo azul infinito. Relájate en nuestras confortables habitaciones, deleita tu paladar con deliciosos sabores locales y descubre un oasis de descanso y rejuvenecimiento en nuestro spa de clase mundial.",
		},
		{
			Nombre:      "Sunset Paradise",
			CantHab:     10,
			Descripcion: "Bienvenido a nuestro acogedor hotel de montaña, rodeado de majestuosas cumbres y bosques exuberantes. Descubre el encanto rústico de nuestras habitaciones, deleita tus sentidos con la cocina alpina en nuestro restaurante y explora las emocionantes actividades al aire libre que te esperan justo afuera de nuestras puertas.",
		},
		{
			Nombre:      "Golden Sands Resort",
			CantHab:     10,
			Descripcion: "Un refugio urbano de estilo contemporáneo, ubicado en el corazón vibrante de la ciudad. Nuestras habitaciones modernas y elegantes te brindarán el máximo confort, mientras que nuestra ubicación privilegiada te permitirá explorar los icónicos lugares de interés, disfrutar de la animada vida nocturna y sumergirte en la cultura local.",
		},
		{
			Nombre:      "Ocean View Inn",
			CantHab:     10,
			Descripcion: "Descubre la elegancia clásica de nuestro hotel histórico, donde la arquitectura impresionante se combina con el encanto del viejo mundo. Admira la belleza de nuestras habitaciones meticulosamente restauradas, deleita tu paladar con exquisiteces culinarias en nuestro restaurante gourmet y déjate envolver por la atmósfera de sofisticación y tradición que te espera.",
		},
	}

	for _, hotel := range hotels {
		if err := db.Create(&hotel).Error; err != nil {
			log.Error("Failed to insert hotel:", err.Error())
		}
	}

	log.Info("Initial values inserted")
}

func init() {
	// DB Connections Paramters
	DBName := "pruebaHash"
	DBUser := "root"
	DBPass := "arquisoft1"
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	hotelClient.Db = db
	reservaClient.Db = db
	amenitieClient.Db = db
	amenitieHotelClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reservas{})
	db.AutoMigrate(&model.Hotels{})
	db.AutoMigrate(&model.Amenities{})
	db.AutoMigrate(&model.AmenitieHotel{})

	insertInitialData()

	log.Info("Finishing Migration Database Tables")
}
