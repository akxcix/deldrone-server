package mysql

import (
	"database/sql"

	"github.com/deldrone/server/pkg/models"
)

// ListingModel wraps a connection pool
type ListingModel struct {
	DB *sql.DB
}

// Insert inserts a new listing into the database
func (m *ListingModel) Insert(vendorID, price int, description, name string) error {
	stmt := `INSERT INTO listings (vendor_id, listing_price, listing_desc, listing_name)
	VALUES (?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, vendorID, price, description, name)
	if err != nil {
		return err
	}

	return nil
}

// All gets all listing for a particular vendor
func (m *ListingModel) All(vendorID int) ([]*models.Listing, error) {
	stmt := `SELECT list_id, vendor_id, listing_price, listing_desc, listing_name
	FROM listings
	WHERE vendor_id = ?`

	rows, err := m.DB.Query(stmt, vendorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listings := []*models.Listing{}

	for rows.Next() {
		l := &models.Listing{}
		err = rows.Scan(&l.ID, &l.VendorID, &l.Price, &l.Description, &l.Name)
		if err != nil {
			return nil, err
		}

		listings = append(listings, l)
	}

	return listings, nil
}

// Get fetches a listing item
func (m *ListingModel) Get(id int) (*models.Listing, error) {
	stmt := `SELECT list_id, vendor_id, listing_price, listing_desc, listing_name
	FROM listings
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)
	l := &models.Listing{}
	err := row.Scan(&l.ID, &l.VendorID, &l.Price, &l.Description, &l.Name)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return l, nil
}
