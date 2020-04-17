package views

import (
	"fmt"
	"net/http"
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
	fmt.Fprint(w, "GET Request to users")
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
