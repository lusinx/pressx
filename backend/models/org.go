package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)
// Org structure
type Org struct {
	gorm.Model
	// data relating to an organisation
	ID    string `gorm:"unique" gorm:"primary_key"`
	Name  string
	Img   string
	Users []*User
}

// Create method for Org
func (org *Org) Create() (*Org, error) {
	if !db.NewRecord(*org) {
		return nil, fmt.Errorf("attempting to make duplicate Org entry")
	}
	if err := db.Create(org).Error; err != nil {
		return nil, err
	}
	return org, nil
}
