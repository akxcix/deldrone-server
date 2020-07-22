// routes.go contains the routes which the web application will handle
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes defines a router through which we register our handler functions with specific routes
func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(app.default404)

	// TODO: cleanup middleware
	r.Handle("/", noSurf(http.HandlerFunc(app.home))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signupForm))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signup))).Methods("POST")
	r.Handle("/login", noSurf(http.HandlerFunc(app.loginForm))).Methods("GET")
	r.Handle("/login", noSurf(http.HandlerFunc(app.login))).Methods("POST")
	r.Handle("/customer/home", app.requireAuthenticatedCustomer(http.HandlerFunc(app.customerHome))).Methods("GET")
	r.Handle("/vendor/home", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorHome))).Methods("GET")
	r.Handle("/vendor/listings", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorListings))).Methods("GET")
	r.Handle("/listing/create", app.requireAuthenticatedVendor(http.HandlerFunc(app.listingCreateForm))).Methods("GET")
	r.Handle("/listing/create", app.requireAuthenticatedVendor(http.HandlerFunc(app.listingCreate))).Methods("POST")
	r.Handle("/vendor/orders", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorOrders))).Methods("GET")
	r.Handle("/vendor/{vendorID}", app.requireAuthenticatedCustomer(http.HandlerFunc(app.vendorIDPage))).Methods("GET")
	r.Handle("/listing/{listingID}", app.requireAuthenticatedUser(http.HandlerFunc(app.listingID))).Methods("GET")
	r.Handle("/logout", app.requireAuthenticatedUser(http.HandlerFunc(app.logout))).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET") // static files
	return secureHeaders(app.logRequest(r))
}
