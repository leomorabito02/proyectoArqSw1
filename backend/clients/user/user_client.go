package user

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"repo/model"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByEmail(Email string) model.User {
	var user model.User

	Db.Where("email = ?", Email).First(&user)
	log.Debug("User: ", user)

	return user
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TO DO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
