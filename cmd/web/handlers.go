package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deldrone/server/pkg/forms"
	"github.com/deldrone/server/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", nil)
}

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	// validation checks
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password", "phone", "address", "pincode")
	form.MinLength("password", 8)
	form.MatchesPattern("email", forms.RxEmail)
	pincode, err := strconv.Atoi(form.Get("pincode"))
	if err != nil {
		form.Errors.Add("pincode", "enter valid pincode")
	}

	if form.Get("accType") == "vendor" {
		form.Required("gps_lat", "gps_long")
		gpsLat, err := strconv.ParseFloat(form.Get("gps_lat"), 64)
		if err != nil {
			form.Errors.Add("gps_lat", "enter valid value")
		}
		gpsLong, err := strconv.ParseFloat(form.Get("gps_long"), 64)
		if err != nil {
			form.Errors.Add("gps_lat", "enter valid value")
		}
		if !form.Valid() {
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		}

		err = app.vendors.Insert(
			form.Get("name"),
			form.Get("address"),
			form.Get("email"),
			form.Get("password"),
			form.Get("phone"),
			pincode,
			gpsLat,
			gpsLong,
		)
		if err == models.ErrDuplicateEmail {
			form.Errors.Add("email", "Address already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}

	} else {
		if !form.Valid() {
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		}

		err = app.customers.Insert(
			form.Get("name"),
			form.Get("address"),
			form.Get("email"),
			form.Get("password"),
			form.Get("phone"),
			pincode,
		)

		if err == models.ErrDuplicateEmail {
			form.Errors.Add("email", "Address already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}
	}
	w.Write([]byte("signup succesful"))
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.MatchesPattern("email", forms.RxEmail)
	if !form.Valid() {
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	}
	var id int
	if form.Get("accType") == "customer" {
		id, err = app.customers.Authenticate(form.Get("email"), form.Get("password"))
	} else {
		id, err = app.vendors.Authenticate(form.Get("email"), form.Get("password"))
	}
	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Email or Password Incorrect. Please verify you have selected correct account type")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	str := fmt.Sprintf("Login Succesful. ID: %d", id)
	w.Write([]byte(str))
}
