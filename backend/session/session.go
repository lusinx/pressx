package session

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Session struct {
	jwt.StandardClaims
	Username string `json: "username"`
}

func CheckStatus() bool {

}
