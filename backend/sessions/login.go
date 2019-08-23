package sessions

import (
	"net/http"
	"math/rand"
)

// Login is a handler that sets a logged in cookie

type Login struct {
	Name	string
	Value	string
	Path	string
	Domain 	string
	MaxAge  int
	Next http.Handler
}

func GenerateKey(l Login){
	l.Name = "session"
	l.Value = RandStringBytes(32)
	l.Path = "/"
	l.Domain = "localhost"
}

func (l Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie {
		Name:	l.Name,
		Value:	l.Value,
		Domain: l.Domain,
		Path: 	l.Path,
		MaxAge: l.MaxAge,
		HttpOnly: true,
	}

	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	randString := make([]byte, n)
	for i := range randString {
		randString[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(randString)
}