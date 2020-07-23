package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamadarshk/deldrone-server/pkg/forms"
	"github.com/iamadarshk/deldrone-server/pkg/models"
)

// Not Found
func (app *application) default404(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "404.page.tmpl", nil)
}

// Home -------------------------------------------------------------------------------------------

// home shows a home page according to login status
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if app.authenticatedCustomer(r) != 0 {
		http.Redirect(w, r, "/customer/home", http.StatusFound)
		return
	} else if app.authenticatedVendor(r) != 0 {
		http.Redirect(w, r, "/vendor/home", http.StatusFound)
		return
	}
	app.render(w, r, "home.page.tmpl", nil)
}

// SignUp -----------------------------------------------------------------------------------------

// signupForm shows a form for users to signup
func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// signup handles the signup process.
// it validates the form, creates users and handles related errors
func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	// get session
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
	}

	err = r.ParseForm()
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

	// check whether signup was as a vendor or a customer
	if form.Get("accType") == "vendor" {

		// additional GPS validations
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

		// insert into database
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

		// return if duplicate email or some other error
		if err == models.ErrDuplicateEmail {
			form.Errors.Add("email", "Address already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}

	} else { // customer
		// if validation checks didn't pass
		if !form.Valid() {
			// prompt user to fill form with correct data
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		}

		// insert into database
		err = app.customers.Insert(
			form.Get("name"),
			form.Get("address"),
			form.Get("email"),
			form.Get("password"),
			form.Get("phone"),
			pincode,
		)

		// return if duplicate email or some other error
		if err == models.ErrDuplicateEmail {
			form.Errors.Add("email", "Address already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}
	}

	// redirect after succesful signup
	session.AddFlash("Sign Up succesful")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Login ------------------------------------------------------------------------------------------

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = r.ParseForm()
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
		if err == models.ErrInvalidCredentials {
			form.Errors.Add("generic", "Email or Password Incorrect. Please ensure you have selected correct account type")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}
		session.Values["customerID"] = id
		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, err)
			return
		}
		http.Redirect(w, r, "/customer/home", http.StatusSeeOther)
	} else {
		id, err = app.vendors.Authenticate(form.Get("email"), form.Get("password"))
		if err == models.ErrInvalidCredentials {
			form.Errors.Add("generic", "Email or Password Incorrect. Please ensure you have selected correct account type")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}
		session.Values["vendorID"] = id
		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, err)
			return
		}
		http.Redirect(w, r, "/vendor/home", http.StatusSeeOther)
	}
}

// Logout -----------------------------------------------------------------------------------------

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
		return
	}

	if session.Values["customerID"] != nil {
		session.Values["customerID"] = nil
		session.AddFlash("customer logged out")
		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if session.Values["vendorID"] != nil {
		session.Values["vendorID"] = nil
		session.AddFlash("vendor logged out")
		err = session.Save(r, w)
		if err != nil {
			app.serverError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Vendor -----------------------------------------------------------------------------------------

func (app *application) vendorHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "vendorhome.page.tmpl", nil)
}

func (app *application) vendorListings(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
		return
	}
	vendorID := session.Values["vendorID"].(int)

	listings, err := app.listings.All(vendorID)
	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, r, "vendorlistings.page.tmpl", &templateData{Listings: listings})
}

func (app *application) vendorOrders(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "vendororders.page.tmpl", nil)
}

// Customer ---------------------------------------------------------------------------------------

func (app *application) customerHome(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
		return
	}
	customerID := session.Values["customerID"].(int)
	customer, err := app.customers.Get(customerID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	vendors, err := app.vendors.GetByPincode(customer.Pincode, 5)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "customerhome.page.tmpl", &templateData{Vendors: vendors})
}

func (app *application) vendorIDPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID, err := strconv.Atoi(vars["vendorID"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	vendor, err := app.vendors.Get(vendorID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	listings, err := app.listings.All(vendorID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "vendorid.page.tmpl", &templateData{
		Vendor:   vendor,
		Listings: listings,
	})
}

func (app *application) customerCart(w http.ResponseWriter, r *http.Request) {
	customerID := app.authenticatedCustomer(r)
	cart := app.carts[customerID]

	cartRowSlice := make([]cartRow, 0)
	total := 0

	for listID, quantity := range cart {
		listing, err := app.listings.Get(listID)
		if err != nil {
			app.serverError(w, err)
		}
		row := cartRow{
			ListingID: listing.ID,
			Name:      listing.Name,
			Price:     listing.Price,
			Quantity:  quantity,
			Amount:    quantity * listing.Price,
		}
		cartRowSlice = append(cartRowSlice, row)
		total += quantity * listing.Price
	}
	app.render(w, r, "customercart.page.tmpl", &templateData{
		Cart:      cartRowSlice,
		CartTotal: total,
	})
	return
}

func (app *application) customerAddToCart(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
	}
	vars := mux.Vars(r)
	listingID, err := strconv.Atoi(vars["listingID"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("quantity")
	quantity, err := strconv.Atoi(form.Get("quantity"))
	if err != nil || quantity <= 0 {
		form.Errors.Add("quantity", "Quantity must be a positive integer.")
	}
	if !form.Valid() {
		listing, err := app.listings.Get(listingID)
		if err == models.ErrNoRecord {
			app.clientError(w, http.StatusBadRequest)
			return
		} else if err != nil {
			app.serverError(w, err)
		}

		app.render(w, r, "listingid.page.tmpl", &templateData{
			Listing: listing,
			Form:    form,
		})
		return
	}

	customerID := app.authenticatedCustomer(r)
	cart := app.carts[customerID]
	if cart == nil {
		cart = models.Cart{}
	}
	app.carts[customerID] = cart.Add(listingID, quantity)
	session.AddFlash("Added to Cart")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}
	url := fmt.Sprintf("/listing/%d", listingID)
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

func (app *application) checkout(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "checkout.page.tmpl", nil)
}

// Listings ---------------------------------------------------------------------------------------
func (app *application) listingCreateForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "listingcreate.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) listingCreate(w http.ResponseWriter, r *http.Request) {
	session, err := app.sessionStore.Get(r, "session-name")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "description", "price")
	if !form.Valid() {
		app.render(w, r, "listingcreate.page.tmpl", &templateData{Form: form})
		return
	}
	vendorID := app.authenticatedVendor(r)
	price, err := strconv.Atoi(form.Get("price"))
	if err != nil {
		form.Errors.Add("price", "enter valid integer")
	}
	err = app.listings.Insert(
		vendorID,
		price,
		form.Get("description"),
		form.Get("name"),
	)
	if err != nil {
		app.serverError(w, err)
	}

	session.AddFlash("Succesful Listed")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
	}
	http.Redirect(w, r, "/vendor/listings", http.StatusSeeOther)
}

func (app *application) listingID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listingID, err := strconv.Atoi(vars["listingID"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	listing, err := app.listings.Get(listingID)
	if err == models.ErrNoRecord {
		app.clientError(w, http.StatusBadRequest)
		return
	} else if err != nil {
		app.serverError(w, err)
	}

	app.render(w, r, "listingid.page.tmpl", &templateData{
		Listing: listing,
		Form:    forms.New(nil),
	})
}
