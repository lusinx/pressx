package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Content structure
type Content struct {
	gorm.Model
	//  data contained within content
	format int
	body   string
	user   *User
	id     string
}

// Create method for Content
func (content *Content) Create() (*Content, error) {
	if !db.NewRecord(*content) {
		return nil, fmt.Errorf("attempting to make duplicate entry")
	}
	if err := db.Create(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}
