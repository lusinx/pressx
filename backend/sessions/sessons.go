package sessions

import "github.com/gorilla/sessions"

type Session struct {
	key   [32]byte
	store *sessions.CookieStore
}

func (s *Session) Login() bool {
	return false
}

func init() {
	session := Session{}
}
