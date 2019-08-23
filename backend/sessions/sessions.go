package sessions

import (
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Session struct {
	key   [32]byte
	store *sessions.CookieStore
}

func (s *Session) Login() bool {
	return false
}

func init() {
	session := Session{}

	authKey := securecookie.GenerateRandomKey(32)
	copy(session.key[:], authKey)

	session.store.Options = &sessions.Options{
		HttpOnly: true,
	}
	gob.Register(Session{})
}
