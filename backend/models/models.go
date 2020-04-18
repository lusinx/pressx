package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var db *gorm.DB

// Models methods
type Models interface {
	Create() (interface{}, error)
}

// DBMigrate structure
func DBMigrate() {
	var err error
	if db, err = gorm.Open("postgres", viper.GetString("postgres.config")); err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Org{})
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Content{})
}
