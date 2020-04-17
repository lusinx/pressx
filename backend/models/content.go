package models

// Content structure
type Content struct {
	//  data contained within content
	format int
	body   string
	user   *User
	id     string
}
