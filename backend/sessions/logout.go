package sessions

import(
	"net/http"
)

// Logout is a handler that clears a logged-in cookie

type Logout struct {
	Name	string
	Path	string
	Domain 	string
	Next	http.Handler
}

func (l Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:       l.Name,
		Value:      "",
		Path:       l.Path,
		Domain:     l.Domain,
		//Expires:    time.Time{},
		//RawExpires: "",
		MaxAge:     0,
		//Secure:     false,
		HttpOnly:   false,
		//SameSite:   0,
		//Raw:        "",
		//Unparsed:   nil,
	}

	http.SetCookie(rw, &cookie)
	l.Next.ServeHTTP(rw, r)
}