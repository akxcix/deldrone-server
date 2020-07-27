package mysql

import (
	"database/sql"
	"time"

	"github.com/iamadarshk/deldrone-server/pkg/models"
)

// DeliveryModel wraps a connection pool
type DeliveryModel struct {
	DB *sql.DB
}

// Insert inserts a value into the table
func (m *DeliveryModel) Insert(customerID, vendorID int, placementTime time.Time, dropLat, dropLong float64) error {
	stmt := `INSERT INTO deliveries (cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, customerID, vendorID, placementTime, dropLat, dropLong, "placed")
	return err
}

// Get fetches a delivery
func (m *DeliveryModel) Get(deliveryID int) (*models.Delivery, error) {
	stmt := `SELECT delivery_id, cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status
	FROM deliveries
	WHERE delivery_id = ?`

	row := m.DB.QueryRow(stmt, deliveryID)
	d := &models.Delivery{}
	err := row.Scan(&d.ID, &d.CustomerID, &d.VendorID, &d.TimeOfDelivery, &d.DropLat, &d.DropLong, &d.Status)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return d, nil
}
