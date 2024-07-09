package handlers

import (
	"fmt"
	"main/data"
	"net/http"
)

func LoginProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		value := data.VerifyLogin(username, password)
		if value {
			fmt.Println("login aceito")
			http.Redirect(w, r, "/login/succeeded", http.StatusSeeOther)
		} else {
			fmt.Println("login negado")
			http.Redirect(w, r, "/login/notSucceeded", http.StatusSeeOther)
		}
		fmt.Fprint(w, "ok")
	}
}

func RegisterProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		email := r.FormValue("e-mail")

		value := data.InsertUser(username, email, phone, password)

		if value {
			http.Redirect(w, r, "/register/succeeded", http.StatusSeeOther)
			fmt.Println("registrando")
		} else {
			http.Redirect(w, r, "/register/notSucceeded", http.StatusSeeOther)
			fmt.Println("ja registrado")
		}
	}
}
