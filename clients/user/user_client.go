package clients

package clients

import (
"HotelArquiSoft/model"
log "github.com/sirupsen/logrus"
)

//var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id=?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.User {
	var user model.User
	Db.Find(&user)
	log.Debug("User: ", user)
	return user
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
