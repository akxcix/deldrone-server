package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/deldrone/server/pkg/forms"

	"golang.org/x/crypto/bcrypt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl")
}

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/signup.page.tmpl",
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

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password", "phone", "address", "pincode")
	form.MinLength("password", 8)
	form.MatchesPattern("email", forms.RxEmail)
	if form.Get("accType") == "vendor" {
		form.Required("gps_lat", "gps_long")
		if !form.Valid() {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		w.Write([]byte("Vendor SignUp"))
	} else {
		w.Write([]byte("Customer signup"))
	}
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/login.page.tmpl",
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

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	accType := r.PostForm.Get("accType")
	storedHash := []byte("$2a$12$wryhNqW9750Nd1ZWRZd4leov8.SbD/dUeD13KOhJZhC/86CQ1vSEq")

	err = bcrypt.CompareHashAndPassword(storedHash, []byte(password))
	same := "yes"
	if err == bcrypt.ErrMismatchedHashAndPassword {
		same = "no"
	} else if err != nil {
		same = err.Error()
	}

	str := fmt.Sprintf("%s\n%s\n%s\n%s", email, hashedPassword, accType, same)
	w.Write([]byte(str))

}
