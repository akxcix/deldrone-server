package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is used when no record is found in the database
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials is used when a user provides invalid credentials
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail is used when the email used to sign up already exists in the database
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Customer a model used to interface with customers table in our database
type Customer struct {
	ID             int    `json:"customerID"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Pincode        int    `json:"pincode"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"-"`
}

// Vendor is a model used to interface with the vendors table in our database
type Vendor struct {
	ID             int     `json:"vendorID"`
	Name           string  `json:"name"`
	Pincode        int     `json:"pincode"`
	GpsLat         float64 `json:"latitude"`
	GpsLong        float64 `json:"longitude"`
	Email          string  `json:"email"`
	HashedPassword []byte  `json:"-"`
	Address        string  `json:"address"`
	Phone          int     `json:"phone"`
}

// Listing is a model that interfaces with the listings table in the database
type Listing struct {
	ID          int    `json:"listingID"`
	VendorID    int    `json:"vendorID"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Delivery is a model that interfaces with the deliveries table in the database
type Delivery struct {
	ID             int       `json:"deliveryID"`
	CustomerID     int       `json:"customerID"`
	VendorID       int       `json:"vendorID"`
	TimeOfDelivery time.Time `json:"timeOfDelivery"`
	DropLat        float64   `json:"dropLatitude"`
	DropLong       float64   `json:"dropLongitude"`
	Status         string    `json:"status"`
}

// Order is a model that interfaces with the orders table in the database
type Order struct {
	ID         int `json:"orderID"`
	DeliveryID int `json:"deliveryID"`
	ListingID  int `json:"listingID"`
	Quant      int `json:"quantity"`
	Amount     int `json:"amount"`
}

// Cart is a struct that denotes a user's cart. The Key is the listing's id and value is it's quantity
type Cart map[int]int

// Add adds the item with listingID and quantity to cart
func (cart Cart) Add(listingID, quantity int) Cart {
	cart[listingID] = quantity
	return cart
}

// Get fetches the quantity of an item in cart, returns 0 if not added
func (cart Cart) Get(listingID int) int {
	quantity := cart[listingID]
	return quantity
}
