package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lusinx/pressx/database"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login sets the JWT session token
func Login(w http.ResponseWriter, r *http.Request) {
	var cred *Credentials

	if err := json.NewDecoder(r.Body).Decode(cred); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !database.QueryAuth(cred.Username, cred.Password) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sToken, err := token.SignedString("secret_placeholder")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", fmt.Sprintf("BEARER %s", sToken))
}
