package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deldrone/server/pkg/forms"
	"github.com/deldrone/server/pkg/models"
	"golang.org/x/crypto/bcrypt"
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
	pincodeInt, err := strconv.Atoi(form.Get("pincode"))
	if err != nil {
		form.Errors.Add("pincode", "enter valid pincode")
	}

	if form.Get("accType") == "vendor" {
		form.Required("gps_lat", "gps_long")
		if !form.Valid() {
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		}
		w.Write([]byte("Vendor SignUp"))
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
			pincodeInt,
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
	app.render(w, r, "login.page.tmpl", nil)
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
