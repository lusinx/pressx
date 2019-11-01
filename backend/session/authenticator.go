package session

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lusinx/pressx/database"
)

// Used to Authenticate Moderators with a moderation code
func ModAuthenticator(route *string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, claims, err := jwtauth.FromContext(r.Context())

			if err != nil {
				fmt.Println("Invalid header")
				http.Error(w, http.StatusText(401), 401)
				return
			}

			if token == nil || !token.Valid {
				fmt.Println("Invalid token")
				http.Error(w, http.StatusText(401), 401)
				return
			}

			// Get auth group code, check if valid, then verify
			if code, ok := claims["auth_group"].(float64); ok && database.VerifyAuth(int(code), *route) {
				// Private route
				// Token is authenticated, pass it through
				fmt.Printf("Route authenticated [%s] with auth_group [%d]\n", *route, int(code))
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, http.StatusText(401), 401)

		})
	}
}
