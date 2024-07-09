package main

import (
	"fmt"
	"main/handlers"
	"main/views"
	"main/views/layout"
	"main/views/partial"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	login := layout.Base(views.Login())

	http.Handle("/", templ.Handler(login))
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(layout.Base(views.Register())))
	http.HandleFunc("/login/processing", handlers.LoginProcessing)
	http.Handle("/register/processing", templ.Handler(partial.LoginVerification(true)))

	fmt.Println("Server on at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
