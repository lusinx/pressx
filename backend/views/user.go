package views

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lusinx/pressx/auth"
	"net/http"
	"encoding/json"

	"github.com/lusinx/pressx/models"
	"github.com/gorilla/mux"
)

func parseForm(field string, out *string, missing *[]string, r *http.Request) {
	if ent := r.PostFormValue(field); len(ent) > 0 {
		*out = ent
	} else {
		*missing = append(*missing, field)
	}
}

func checkMissing(missing []string, w *http.ResponseWriter) bool {
	if len(missing) > 0 {
		var so string
		for _, s := range missing {
			so += fmt.Sprintf("%s ", s)
		}
		http.Error(*w, fmt.Sprintf("Missing form fields: %s\n", so), 403)
		return true
	}
	return false
}

// GetUser GET Request
func GetUser(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	var user = mux.Vars(r)["user"]

	w.Write([]byte(user))

	//w.Write([]byte(models.User{user, "Thomas", "Galligan",}))

	// make request with gorm to request user details
	//fmt.Fprint(w, "GET Request to users")
}

// NewUser POST Request
func NewUser(w http.ResponseWriter, r *http.Request) {
	var (
		username,
		password,
		email,
		first,
		last string
		missing []string
	)
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Request malformed: %v", err), 500)
		return
	}
	parseForm("username", &username, &missing, r)
	parseForm("password", &password, &missing, r)
	parseForm("email", &email, &missing, r)
	parseForm("firstname", &first, &missing, r)
	parseForm("lastname", &last, &missing, r)

	if checkMissing(missing, &w) {

		return
	}

	// Generate the user
	salt := auth.GenerateSalt()
	password = auth.HashPassword(password, salt)
	user := models.User{
		Model:      gorm.Model{},
		Username:   "",
		Firstnames: "",
		Lastname:   "",
		Img:        "",
		Perms:      0,
		ViewPerms:  0,
		Orgs:       nil,
		AuthGroup:  0,
	}
	_, err = user.Create()
	if err != nil {
		return
	}
	
	user := models.UserAuth{}

	fmt.Println(user)
	
	user.Create()



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
