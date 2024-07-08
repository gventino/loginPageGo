package main

import (
	"fmt"
	"main/views"
	"main/views/layout"
	"main/views/partial"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	login := layout.Base(views.Login())

	http.Handle("/", templ.Handler(login))
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(layout.Base(views.Register())))
	http.Handle("/verification", templ.Handler(partial.LoginVerification(true)))

	fmt.Println("Server on at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
