package models

import "github.com/jinzhu/gorm"

// UserAuth structure
type UserAuth struct {
	gorm.Model
	Username string
	Password string
	Salt     string
	Perms    uint8
}
