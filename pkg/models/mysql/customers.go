package mysql

import (
	"database/sql"

	"github.com/deldrone/server/pkg/models"
)

// CustomerModel wraps a connection pool
type CustomerModel struct {
	DB *sql.DB
}

// Insert creates a new user by inserting values into the database. Returns an error
func (m *CustomerModel) Insert(name, address, pincode, email, password string, phone int) error {
	return nil
}

// Authenticate verifies the credentials and returns userid if valid details are provided.
func (m *CustomerModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get fetches the details of the customer using its id
func (m *CustomerModel) Get(id int) *models.Customer {
	return nil
}
