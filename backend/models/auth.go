package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// UserAuth structure
type UserAuth struct {
	gorm.Model
	Username string
	Password string
	Salt     string
	Perms    uint8
}

// Create method for Org
func (userauth *UserAuth) Create() (*UserAuth, error) {
	if !db.NewRecord(*userauth) {
		return nil, fmt.Errorf("attempting to make duplicate entry")
	}
	if err := db.Create(userauth).Error; err != nil {
		return nil, err
	}
	return userauth, nil
}
