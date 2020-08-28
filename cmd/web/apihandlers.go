// this file contains the handlers required for serving the api
// TODO: Add Authentication
package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Vendor -----------------------------------------------------------------------------------------

// Method: GET, Path: "/api/vendor/{vendorID}"
// Feteches a particular vendor by their vendorid
func (app *application) apiGetVendorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorid, err := strconv.Atoi(vars["vendorID"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	vendor, err := app.vendors.Get(vendorid)
	if err != nil {
		app.serverError(w, err)
		return
	}
	jsonData, err := json.Marshal(vendor)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.writeJSON(w, r, jsonData)
}

// Method: GET, PATH: "/api/vendors/{pincode}/{pincoderange}"
// Fetches all vendors whith pincode = {pincode} +- {pincoderange}
func (app *application) apiGetVendorByPincode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pincode, err := strconv.Atoi(vars["pincode"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	pincodeRange, err := strconv.Atoi(vars["pincoderange"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	vendors, err := app.vendors.GetByPincode(pincode, pincodeRange)
	if err != nil {
		app.serverError(w, err)
		return
	}
	jsonData, err := json.Marshal(vendors)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.writeJSON(w, r, jsonData)
}

// Method: GET, Path: "/api/vendor/{vendorID}/listings"
// Fetches all listings by particular vendor using vendorID

// Method: PUT, Path: "/api/vendor/{vendorID}/listing/{listingID}"
// Updates particular listing using listingID and vendorID

// Method: DELETE, Path: "/api/vendor/{vendorID}/listing/{listingID}"
// Deletes particular listing using vendorID and listingID

// Method: POST, Path: "/api/vendor/{vendorID}/listing/new"
// Creates a new listing for vendor with their vendorID

// Method: GET, Path: "/api/vendor/{vendorID}/orders"
// fetches all orders from vendor with their vendorID

// Method: GET, Path: "/api/vendor/{vendorID}/activeorders"
// fetches activer orders for vendor using their vendorID

// Customer ---------------------------------------------------------------------------------------

// Method: GET, Path: "api/customer/{customerID}"
// fetches details for particular customer using their customerID

// Method: GET, Path: "api/customer/{customerID}/getcart"
// fetches customer's cart using their customerID

// Method: POST, Path: "api/customer/{customerID}/cart/{listingID}"
// adds listing with listingID in customer's cart using their customerID

// Method: PUT, Path: "api/customer/{customerID}/cart/{listingID}"
// updates listing with listingID in customer's cart using their customerID

// Method: DELETE, Path: "api/customer/{customerID}/cart/{listingID}"
// deletes listing with listingID in customer's cart using their customerID

// Method: POST, Path: "api/customer/{customerID}/checkout"
// checks out customer using their customerID

// Method: GET, Path: "api/customer/{customerID}/orders"
// fetches all orders for customer using their customerID

// Method: GET, Path: "api/customer/{customerID}/activeorders"
// fetches active orders for customer using their customerID

// Listings ---------------------------------------------------------------------------------------

// Method: GET, Path: "api/listing/{listingID}"
// Fetches a listing by it's listingID
func (app *application) apiGetListingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listingID, err := strconv.Atoi(vars["listingID"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	listing, err := app.listings.Get(listingID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	jsonData, err := json.Marshal(listing)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.writeJSON(w, r, jsonData)
}

// Deliveries -------------------------------------------------------------------------------------

// Orders -----------------------------------------------------------------------------------------
