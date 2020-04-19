package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// var db = *gorm.DB

// User structure
type User struct {
	gorm.Model
	// ID will be taken as primary field by default,
	Username   string
	Firstnames string
	Lastname   string
	Img        string
	Perms      uint8 // For frontend accessing endpoints
	ViewPerms  uint8 // For people wanting to view your profile
	Orgs       []*Org
	AuthGroup  int // Admin panel auth
}

// Create method for User
func (user *User) Create() (*User, error) {
	if !db.NewRecord(*user) {
		return nil, fmt.Errorf("attempting to make duplicate User entry")
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func Search(username string) (*User, error) {

	user := User{}

	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	if user.ViewPerms >= 3 {
		return nil, nil
	}

	return &user, nil
}
