package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// ID will be taken as primary field by default,
	ID              string	`gorm:"unique" gorm:"primary_key"`
	Username        string	
	Img             string	
	Permissions     uint8 	// For frontend auth
	ViewPermissions uint8	
	Orgs            []*Org	
	AuthGroup       int 	// Admin panel auth
}

type Org struct {
	// data relating to an organisation
	gorm.Model
	ID    string	`gorm:"unique" gorm:"primary_key"`
	Name  string
	Img   string
	Users []*User
}

type Page struct {
	// data contained within a page
	Name string
	//content  datatype
}

type Image struct {
	// data contained within an image
	External bool
	Slug     string		`gorm:"unique" gorm:"primary_key"`//where to look for it
	User     *User
	CdnPath  string
}

type Content struct {
	//  data contained within content
	format int
	body   string
	user   *User
	id     string
}


func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Org{})
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Image{})
	
	return db
}