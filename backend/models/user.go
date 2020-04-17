package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User structure
type User struct {
	gorm.Model
	// ID will be taken as primary field by default,
	Username   string
	Firstnames string
	Lastname   string
	Img        string
	Perms      uint8 // For frontend auth
	ViewPerms  uint8
	Orgs       []*Org
	AuthGroup  int // Admin panel auth
}

// Create method for User
func (user *User) Create() (*User, error) {
	if !db.NewRecord(*user) {
		return nil, fmt.Errorf("attempting to make duplicate User entry")
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
