package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// defines a router through which we register our handler functions to routes
func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	return r
}
