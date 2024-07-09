package handlers

import (
	"fmt"
	"net/http"
)

func LoginProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println(username + password)
	}

}
