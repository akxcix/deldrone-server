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

// Customer a model used to interface with custtomers table in our database
type Customer struct {
	ID             int
	Name           string
	Address        string
	Pincode        int
	Phone          string
	Email          string
	HashedPassword []byte
}

// Vendor is a model used to interface with the vendors table in our database
type Vendor struct {
	ID             int
	Name           string
	Pincode        int
	GpsLat         float64
	GpsLong        float64
	Email          string
	HashedPassword []byte
	Address        string
	Phone          int
}

// Listing is a model that interfaces with the listings table in the database
type Listing struct {
	ID          int
	VendorID    int
	Price       int
	Description string
	Name        string
}

// Delivery is a model that interfaces with the deliveries table in the database
type Delivery struct {
	ID             int
	CustomerID     int
	VendorID       int
	TimeOfDelivery time.Time
	DropLat        float64
	DropLong       float64
	Status         string
}

// Order is a model that interfaces with the orders table in the database
type Order struct {
	ID         int
	DeliveryID int
	ListingID  int
	Quant      int
	Amount     int
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
