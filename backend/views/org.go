package views

import (
	"fmt"
	"net/http"
)

// GetOrg GET Request
func GetOrg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET request to Org")
}

// NewOrg POST Request
func NewOrg(w http.ResponseWriter, r *http.Request) {

}

// UpdateOrg UPDATE Request
func UpdateOrg(w http.ResponseWriter, r *http.Request) {

}

// DeleteOrg DELETE Request
func DeleteOrg(w http.ResponseWriter, r *http.Request) {

}

//Subdirectory /org/settings

// GetOrgSettings GET Request
func GetOrgSettings(w http.ResponseWriter, r *http.Request) {

}

// UpdateOrgSettings PUT Request
func UpdateOrgSettings(w http.ResponseWriter, r *http.Request) {

}
