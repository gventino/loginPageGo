package handlers

import (
	"crypto/sha256"
	"fmt"
	"main/data"
	"net/http"
)

func LoginProcessing(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		encodedPassword := sha256.Sum256([]byte(password))
		value := data.VerifyLogin(username, string(encodedPassword[:]))
		if value {
			fmt.Println("login aceito")
		} else {
			fmt.Println("login negado")
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

		encodedPassword := sha256.Sum256([]byte(password))
		value := data.InsertUser(username, email, phone, string(string(encodedPassword[:])))

		if value {
			fmt.Println("registrando")
		} else {
			fmt.Println("ja registrado")
		}
	}
}
