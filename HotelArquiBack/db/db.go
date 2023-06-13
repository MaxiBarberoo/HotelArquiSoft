package db

import (
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
			Nombre:  "Luxury",
			CantHab: 10,
		},
		{
			Nombre:  "Grand Hotel",
			CantHab: 10,
		},
		{
			Nombre:  "Sunset Paradise",
			CantHab: 10,
		},
		{
			Nombre:  "Golden Sands Resort",
			CantHab: 10,
		},
		{
			Nombre:  "Ocean View Inn",
			CantHab: 10,
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
	DBName := "HotelArquiSoft"
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

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Reservas{})
	db.AutoMigrate(&model.Hotels{})

	insertInitialData()

	log.Info("Finishing Migration Database Tables")
}
