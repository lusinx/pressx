package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth"
)

// AuthUser returns auth middleware for a JWT
func AuthUser(perms uint8) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
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
			permsClaim, ok := claims["perms"].(float64)
			if ok && uint8(permsClaim) <= perms {
				// Token is authenticated, pass it through
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, http.StatusText(403), 403)

		})
	}
}
