package views

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/lusinx/pressx/models"
)

// These 4 views depend on the session.
// The user has to be signed in to get proper functionality from these views.

func GetUser(db *gorm.DB) {

	db.Find(models.User{})

}

func NewUser(data models.User, db *gorm.DB) {

	db.Save(data)

}

func UpdateUser(data models.User, db *gorm.DB) {

	db.Model(data).Updates(models.User{})

}

func DeleteUser(id string, db *gorm.DB) {

	db.Where("id = %s", id).Delete(models.User{})

}

//Subdirectory /user/settings

// func GetUserSetting(id string, db *gorm.DB) {

// 	// db.Where("id = %s", id).Find(models.Settings{})

// }
// func PatchUserSettings(data models.Settings, db *gorm.DB){


	
// }