package session

import (
	"encoding/json"
	"net/http"

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

}
