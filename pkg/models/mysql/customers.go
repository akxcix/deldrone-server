package mysql

import (
	"database/sql"

	"github.com/deldrone/server/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// CustomerModel wraps a connection pool
type CustomerModel struct {
	DB *sql.DB
}

// Insert creates a new user by inserting values into the database. Returns an error
func (m *CustomerModel) Insert(name, address, email, password, phone string, pincode int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO customers (cust_name, cust_address, cust_pincode, cust_email, cust_hash_pwd, cust_phone)
	VALUES(?, ?, ?, ?, ?, ?)`

	_, err = m.DB.Exec(stmt, name, address, pincode, email, hashedPassword, phone)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

// Authenticate verifies the credentials and returns userid if valid details are provided.
func (m *CustomerModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get fetches the details of the customer using its id
func (m *CustomerModel) Get(id int) *models.Customer {
	return nil
}
