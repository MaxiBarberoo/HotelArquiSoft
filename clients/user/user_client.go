package clients

import (
	"HotelArquiSoft/model"
<<<<<<< HEAD
=======

>>>>>>> 6a081fd08874f7ae249b2cdd28b49cfb09637103
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
<<<<<<< HEAD
	log.Debug("User: ", users)
	return users
=======

	log.Debug("Users: ", users)
>>>>>>> 6a081fd08874f7ae249b2cdd28b49cfb09637103

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("User Created: ", user.ID)
	return user
}
