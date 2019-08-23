package models

type User struct {
	id          string
	username    string
	img 		string
	permissions int8
	orgs        []*Org
}

type Org struct {
	id    string
	name  string
	img   string
	users []*User
}

type Page   struct {
	name    string
	//content  datatype
}

type Image struct {
	external	bool
	slug		string //where to look for it
	user		*User
	cdnPath 	string
}