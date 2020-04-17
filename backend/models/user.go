package models

import "github.com/jinzhu/gorm"

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
