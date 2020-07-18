package models

import "errors"

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
