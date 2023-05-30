package clients

import (
	"HotelArquiSoft/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

)
var Db *gorm.DB

func  GetAdminById(id int) model.Admin{
	var admin model.Admin

	Db.Where("id = ?", id).First.(&admin)
	log.Debug("Admin: ",admin)
	return admin
}


func GetAdmins() model.Admin {
	var admins model.Admin
	Db.Find(&admins)
	log.Debug("Admins: ",admins)
	return admins
}

func InsertAdmin(admin model.Admin) model.Admin{
	result:= Db.Create(&admin)
	if(result.Error != nil){
		log.Error("")
	}
	log.Debug("Admin Created: ",admin.Id)
	return admin
}
