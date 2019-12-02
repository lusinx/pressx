package models

type User struct {
	// data about a user
	id              string
	username        string
	img             string
	permissions     uint8 // For frontend auth
	viewPermissions uint8
	orgs            []*Org
	authGroup       int // Admin panel auth
}

type Org struct {
	// data relating to an organisation
	id    string
	name  string
	img   string
	users []*User
}

type Page struct {
	// data contained in a page
	name    string
	content *Content
}

type Image struct {
	// data contained within an image type
	external bool
	slug     string //where to look for it
	user     *User
	cdnPath  string
}

type Content struct {
	//  data contained within content
	format int
	body   string
	user   *User
	id     string
}
