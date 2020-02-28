package views

import (
	"fmt"
	"net/http"
)

func GetPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get request to Page")
}

func NewPage(w http.ResponseWriter, r *http.Request) {

}

func UpdatePage(w http.ResponseWriter, r *http.Request) {

}

func DeletePage(w http.ResponseWriter, r *http.Request) {

}

//subdirectory /page/{page}
