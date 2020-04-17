package models

// Image structure
type Image struct {
	// data contained within an image
	External bool
	Slug     string `gorm:"unique" gorm:"primary_key"` //where to look for it
	User     *User
	CdnPath  string
}
