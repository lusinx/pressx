package session

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lusinx/pressx/database"
)

// Used to Authenticate Moderators with a moderation code
func ModAuthenticator(route string) func(http.Handler) http.Handler {
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

			// Get auth group code, check if valid, then verify
			if code, ok := claims["auth_group"].(int); ok && database.VerifyAuth(code, route) {
				// Private route
				// Token is authenticated, pass it through
				next.ServeHTTP(w, r)
			}

			http.Error(w, http.StatusText(401), 401)

		})
	}
}
