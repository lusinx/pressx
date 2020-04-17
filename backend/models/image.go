import (
	"fmt"

	"github.com/jinzhu/gorm"
)
// Image structure
type Image struct {
	gorm.Model
	// data contained within an image
	External bool
	Slug     string `gorm:"unique" gorm:"primary_key"` //where to look for it
	User     *User
	CdnPath  string
}

// Create method for Image
func (image *Image) Create() (*Image, error) {
	if !db.NewRecord(*image) {
		return nil, fmt.Errorf("attempting to make duplicate Image entry")
	}
	if err := db.Create(image).Error; err != nil {
		return nil, err
	}
	return image, nil
}
