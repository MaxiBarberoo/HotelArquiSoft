package clients

import (
	"HotelArquiSoft/model"
	log "github.com/sirupsen/logrus"
)

//var Db *gorm.DB

func GetAdminById(id int) model.Admin {
	var admin model.Admin

	Db.Where("id=?", id).First(&admin)
	log.Debug("Admin: ", admin)

	return admin
}

func GetAdmins() model.Admin {
	var admin model.Admin
	Db.Find(&admin)
	log.Debug("Admin: ", admin)
	return admin
}

func InsertAdmin(admin model.Admin) model.Admin {
	result := Db.Create(&admin)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Admin Created: ", admin.Id)
	return admin
}
