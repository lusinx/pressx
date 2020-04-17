package views

import (
	"fmt"
	"net/http"
)

// GetPage GET Request
func GetPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get request to Page")
}

// NewPage POST Request
func NewPage(w http.ResponseWriter, r *http.Request) {

}

// UpdatePage PUT Request
func UpdatePage(w http.ResponseWriter, r *http.Request) {

}

// DeletePage DELETE Request
func DeletePage(w http.ResponseWriter, r *http.Request) {

}

//subdirectory /page/{page}
