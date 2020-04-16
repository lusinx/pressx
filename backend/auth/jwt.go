package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/lusinx/pressx/models"
	"github.com/spf13/viper"
)

// NewJWTInstance makes the JWT instance
func NewJWTInstance() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", viper.Get("jwt.token").([]byte), nil)
}

// UserSession returns a new JWT with claims: {user_id: uint, perms: uint}
func UserSession(instance *jwtauth.JWTAuth, user *models.UserAuth) (*jwt.Token, error) {
	if token, str, err := instance.Encode(jwt.MapClaims{"user_id": user.ID, "perms": user.Perms}); err != nil {
		return nil, err
	} else {
		fmt.Printf("Gen JWT for user %d::\n%s\n", user.ID, str)
		return token, nil
	}
}
