package models

import "github.com/jinzhu/gorm"

// Models methods
type Models interface {
	Create() (interface{}, error)
}

// DBMigrate structure
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Org{})
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Content{})

	return db
}
