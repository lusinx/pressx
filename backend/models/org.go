package models

type Org struct {
	// data relating to an organisation
	ID    string `gorm:"unique" gorm:"primary_key"`
	Name  string
	Img   string
	Users []*User
}
