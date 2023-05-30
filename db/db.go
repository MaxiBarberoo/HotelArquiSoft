package db

import (
	adminClient "HotelArquiSoft/clients/admin"
	userClient "HotelArquiSoft/clients/user"
	"HotelArquiSoft/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "HotelArquiSoft"
	DBUser := "ABGMM"
	DBPass := ""
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
	adminClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admins{})
	db.AutoMigrate(&model.Reservas{})
	db.AutoMigrate(&model.Hotels{})

	log.Info("Finishing Migration Database Tables")
}
