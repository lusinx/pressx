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
	// Sample token for demo
	TokenAuth = jwtauth.New("HS256", "jrhTo4yE5lELmaVpYsq05Jcy2C264M7e", nil)
	// auth_group will be mapped to a table of permissions
	_, str, _ := TokenAuth.Encode(jwtauth.Claims{"auth_group": 0})
	fmt.Printf("JWT generated: %s", str)
}
func CheckStatus() bool {
	return false
}
