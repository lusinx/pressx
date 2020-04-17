package models

import "fmt"

// Page structure
type Page struct {
	// data contained within a page
	Name string
	//content  datatype
	Content *Content
}

// Create method for Page
func (page *Page) Create() (*Page, error) {
	if !db.NewRecord(*page) {
		return nil, fmt.Errorf("attempting to make duplicate Page entry")
	}
	if err := db.Create(page).Error; err != nil {
		return nil, err
	}
	return page, nil
}
