package session

import (
	"net/http"

	"github.com/go-chi/jwtauth"
)

// until you decide to write something similar and customize your client response.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			return
		}

		if token == nil || !token.Valid {
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
