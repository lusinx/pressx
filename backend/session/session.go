package session

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Session struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

func CheckStatus() bool {
	return false
}
