package session

import "net/http"

type Credentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

// Login sets the JWT session token
func Login(w http.ResponseWriter, r *http.Request) {

}
