package models

type User struct {
	id          string
	username    string
	permissions int8
	orgs        []*Org
}

type Org struct {
	id    string
	name  string
	users []*User
}

type Setting interface {
	setting_func()
}

type Settings = []Setting
