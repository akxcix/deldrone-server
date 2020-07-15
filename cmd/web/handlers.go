package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.errorLog.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displays form to sign up"))
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displays form to log in"))
}
