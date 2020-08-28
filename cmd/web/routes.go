// routes.go contains the routes which the web application will handle
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes defines a router through which we register our handler functions with specific routes
func (app *application) routes() http.Handler {
	// TODO: Cleanup middleware
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(app.default404)

	// Account Routes
	r.Handle("/", noSurf(http.HandlerFunc(app.home))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signupForm))).Methods("GET")
	r.Handle("/signup", noSurf(http.HandlerFunc(app.signup))).Methods("POST")
	r.Handle("/login", noSurf(http.HandlerFunc(app.loginForm))).Methods("GET")
	r.Handle("/login", noSurf(http.HandlerFunc(app.login))).Methods("POST")
	r.Handle("/logout", app.requireAuthenticatedUser(http.HandlerFunc(app.logout))).Methods("POST")

	// Customer Routes:
	r.Handle("/customer/home", app.requireAuthenticatedCustomer(http.HandlerFunc(app.customerHome))).Methods("GET")
	r.Handle("/customer/cart", app.requireAuthenticatedCustomer(http.HandlerFunc(app.customerCart))).Methods("GET")
	r.Handle("/customer/checkout", app.requireAuthenticatedCustomer(http.HandlerFunc(app.checkoutForm))).Methods("GET")
	r.Handle("/customer/checkout", app.requireAuthenticatedCustomer(http.HandlerFunc(app.checkout))).Methods("POST")
	r.Handle("/customer/addtocart/{listingID}", app.requireAuthenticatedCustomer(http.HandlerFunc(app.customerAddToCart))).Methods("POST")
	r.Handle("/customer/activeorders", app.requireAuthenticatedCustomer(http.HandlerFunc(app.activeOrders))).Methods("GET")
	r.Handle("/customer/pastorders", app.requireAuthenticatedCustomer(http.HandlerFunc(app.pastOrders))).Methods("GET")

	// Vendor Routes
	r.Handle("/vendor/home", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorHome))).Methods("GET")
	r.Handle("/vendor/listings", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorListings))).Methods("GET")
	r.Handle("/listing/create", app.requireAuthenticatedVendor(http.HandlerFunc(app.listingCreateForm))).Methods("GET")
	r.Handle("/listing/create", app.requireAuthenticatedVendor(http.HandlerFunc(app.listingCreate))).Methods("POST")
	r.Handle("/vendor/orders", app.requireAuthenticatedVendor(http.HandlerFunc(app.vendorOrders))).Methods("GET")
	r.Handle("/vendor/{vendorID}", app.requireAuthenticatedCustomer(http.HandlerFunc(app.vendorIDPage))).Methods("GET")

	// Listing Routes
	r.Handle("/listing/{listingID}", app.requireAuthenticatedUser(http.HandlerFunc(app.listingID))).Methods("GET")

	// Delivery Routes
	r.Handle("/delivery/{deliveryID}", app.requireAuthenticatedUser(http.HandlerFunc(app.deliveryByID))).Methods("GET")

	// Order Routes
	r.Handle("/order/{orderID}", app.requireAuthenticatedUser(http.HandlerFunc(app.orderByID))).Methods("GET")

	// API ----------------------------------------------------------------------------------------
	r.Handle("/api/vendor/{vendorID}", noSurf(http.HandlerFunc(app.apiGetVendorByID)))
	r.Handle("/api/vendors/{pincode}/{pincoderange}", noSurf(http.HandlerFunc(app.apiGetVendorByPincode)))
	r.Handle("/api/listing/{listingID}", noSurf(http.HandlerFunc(app.apiGetListingByID)))

	// Static file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods("GET")

	return secureHeaders(app.logRequest(r))
}
