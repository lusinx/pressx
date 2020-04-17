package views

import (
	"fmt"
	"net/http"
)

// GetUser GET Request
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET Request to users")
}

// NewUser POST Request
func NewUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser PUT Request
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser DELETE Request
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

//Subdirectory /User/settings

// GetUserSettings GET Request
func GetUserSettings(w http.ResponseWriter, r *http.Request) {

}

// UpdateUserSettings PUT Request
func UpdateUserSettings(w http.ResponseWriter, r *http.Request) {

}
