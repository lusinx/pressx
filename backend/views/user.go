package views

import (
	"net/http"
)

// These 4 views depend on the session.
// The user has to be signed in to get proper functionality from these views.

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func NewUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

//Subdirectory /user/settings

func GetUserSettings(w http.ResponseWriter, r *http.Request) {

}
func PatchUserSettings(w http.ResponseWriter, r *http.Request){

}