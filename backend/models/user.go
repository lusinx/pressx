package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// ID will be taken as primary field by default,
	// ID              string `gorm:"unique" gorm:"primary_key"`
	Username        string
	Firstnames      string
	Lastname        string
	Img             string
	Permissions     uint8 // For frontend auth
	ViewPermissions uint8
	Orgs            []*Org
	AuthGroup       int // Admin panel auth
}
