// routes.go contains the routes which the web application will handle
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes defines a router through which we register our handler functions with specific routes
func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/signup", app.signupForm).Methods("GET")
	r.HandleFunc("/login", app.loginForm).Methods("GET")
	r.HandleFunc("/login", app.login).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET") // static files
	return app.logRequest(r)
}
