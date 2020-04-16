package models

import "github.com/jinzhu/gorm"

// UserAuth struct to store items necessary for login
type UserAuth struct {
	gorm.Model
	Username string
	Password string
	Salt     string
	Perms    uint8
}
