package views

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET Request to users")
}
func NewUser(w http.ResponseWriter, r *http.Request) {

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

//Subdirectory /User/settings

func GetUserSettings(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserSettings(w http.ResponseWriter, r *http.Request) {

}
