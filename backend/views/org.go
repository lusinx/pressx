package views

import (
	"fmt"
	"net/http"
)

func GetOrg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET request to Org")
}
func NewOrg(w http.ResponseWriter, r *http.Request) {

}
func UpdateOrg(w http.ResponseWriter, r *http.Request) {

}

func DeleteOrg(w http.ResponseWriter, r *http.Request) {

}

//Subdirectory /org/settings

func GetOrgSettings(w http.ResponseWriter, r *http.Request) {

}

func UpdateOrgSettings(w http.ResponseWriter, r *http.Request) {

}
