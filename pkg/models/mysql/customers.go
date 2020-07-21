package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/iamadarshk/deldrone-server/pkg/models"
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
	var id int
	var hashedPassword []byte
	stmt := `SELECT cust_id, cust_hash_pwd FROM customers WHERE cust_email = ?`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

// Get fetches the details of the customer using its id
func (m *CustomerModel) Get(id int) *models.Customer {
	return nil
}
