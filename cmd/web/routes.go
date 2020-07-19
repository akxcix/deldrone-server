// routes.go contains the routes which the web application will handle
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes defines a router through which we register our handler functions with specific routes
func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	// TODO: cleanup middleware
	r.Handle("/", noSurf(http.HandlerFunc(app.home))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signupForm))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signup))).Methods("POST")
	r.Handle("/login", noSurf(http.HandlerFunc(app.loginForm))).Methods("GET")
	r.Handle("/login", noSurf(http.HandlerFunc(app.login))).Methods("POST")
	r.Handle("/customer/home", app.requireAuthenticatedCustomer(http.HandlerFunc(app.customerHome))).Methods("GET")
	r.Handle("/vendor/home", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorHome))).Methods("GET")
	r.Handle("/logout", app.requireAuthenticatedUser(http.HandlerFunc(app.logout))).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET") // static files
	return secureHeaders(app.logRequest(r))
}
