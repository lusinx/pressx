package session

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func init() {
	TokenAuth = jwtauth.New("HS256", "secret_placeholder", nil)
	// auth_group will be mapped to a table of permissions
	_, str, _ := TokenAuth.Encode(jwtauth.Claims{"auth_group": 0})
	fmt.Printf("JWT generated: %s", str)
}
func CheckStatus() bool {
	return false
}
