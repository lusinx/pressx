package session

import (
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
}
func CheckStatus() bool {
	return false
}
