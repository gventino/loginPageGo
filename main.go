package main

import (
	"fmt"
	"main/data"
	"main/handlers"
	"main/views"
	"main/views/layout"
	"main/views/partial"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	//data.StartDB()
	login := layout.Base(views.Login())

	http.Handle("/", templ.Handler(login))
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(layout.Base(views.Register())))

	http.HandleFunc("/login/processing", handlers.LoginProcessing)
	http.HandleFunc("/register/processing", handlers.RegisterProcessing)

	http.Handle("/login/succeeded", templ.Handler(layout.Base(partial.Verification("login", true))))
	http.Handle("/login/notSucceeded", templ.Handler(layout.Base(partial.Verification("login", false))))
	http.Handle("/register/succeeded", templ.Handler(layout.Base(partial.Verification("register", true))))
	http.Handle("/register/notSucceeded", templ.Handler(layout.Base(partial.Verification("register", false))))

	http.Handle("/admin", templ.Handler(layout.Base(views.Admin())))
	http.HandleFunc("admin/reset", func(w http.ResponseWriter, r *http.Request) {
		data.ResetDB()
		fmt.Fprint(w, "reset ok")
	})

	fmt.Println("Server on at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
