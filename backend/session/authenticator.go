package session

import (
	"net/http"

	"github.com/go-chi/jwtauth"
)

// until you decide to write something similar and customize your client response.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if token == nil || !token.Valid {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if claims["admin"] == true {
			// Private route
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
