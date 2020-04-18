package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lusinx/pressx/auth"
	"github.com/lusinx/pressx/models"
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
	var username = mux.Vars(r)["user"]

	user, err := models.SearchUser(username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if user == nil {
		http.Error(w, "User not found", 404)
	}

	encoded, err := json.Marshal(*user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(w, encoded)

	//w.Write([]byte(user))

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
		w.Write([]byte("Request malformed"))
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

	user := models.User{
		Model:      gorm.Model{},
		Username:   username,
		Firstnames: first,
		Lastname:   last,
		Img:        "",
		Perms:      5,
		ViewPerms:  2,
		Orgs:       nil,
		AuthGroup:  5,
	}

	// Generate the user
	salt := auth.GenerateSalt()
	password = auth.HashPassword(password, salt)

	_, err := user.Create()

	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	fmt.Printf("User %s created", username)
	userAuth := models.UserAuth{
		Model:    gorm.Model{},
		Username: user.Username,
		Password: password,
		Salt:     salt,
		Perms:    user.Perms,
	}

	_, err = userAuth.Create()
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	encoded, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	respondJSON(w, encoded)
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
