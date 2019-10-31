package models

type User struct {
	id              string
	username        string
	img             string
	permissions     uint8 // For frontend auth
	viewPermissions uint8
	orgs            []*Org
	authGroup       int // Admin panel auth
}

type Org struct {
	id    string
	name  string
	img   string
	users []*User
}

type Page struct {
	name string
	//content  datatype
}

type Image struct {
	external bool
	slug     string //where to look for it
	user     *User
	cdnPath  string
}
