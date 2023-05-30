package clients

<<<<<<< HEAD
import (
	"HotelArquiSoft/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB
=======
package clients

import (
"HotelArquiSoft/model"
log "github.com/sirupsen/logrus"
)

//var Db *gorm.DB
>>>>>>> 228bb09fdc9d5c9d646205be451ea9294f245d31

func GetUserById(id int) model.User {
	var user model.User

<<<<<<< HEAD
	Db.Where("id = ?", id).First(&user)
=======
	Db.Where("id=?", id).First(&user)
>>>>>>> 228bb09fdc9d5c9d646205be451ea9294f245d31
	log.Debug("User: ", user)

	return user
}

<<<<<<< HEAD
func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
=======
func GetUsers() model.User {
	var user model.User
	Db.Find(&user)
	log.Debug("User: ", user)
	return user
>>>>>>> 228bb09fdc9d5c9d646205be451ea9294f245d31
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
<<<<<<< HEAD
=======
		//TODO Manage Errors
>>>>>>> 228bb09fdc9d5c9d646205be451ea9294f245d31
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
