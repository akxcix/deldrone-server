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
func (m *DeliveryModel) Insert(customerID, vendorID int, placementTime time.Time, dropLat, dropLong float64) (int, error) {
	stmt := `INSERT INTO deliveries (cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status)
	VALUES (?, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, customerID, vendorID, placementTime, dropLat, dropLong, "placed")
	if err != nil {
		return 0, err
	}
	deliveryID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(deliveryID), err
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

// GetAllByVendorIDStatus fetches all deliveries associated with a vendor id having a particular status
func (m *DeliveryModel) GetAllByVendorIDStatus(vendorID int, status string) ([]*models.Delivery, error) {
	stmt := `SELECT delivery_id, cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status
	FROM deliveries
	WHERE vendor_id = ?
	AND delivery_status = ?`

	rows, err := m.DB.Query(stmt, vendorID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deliveries := []*models.Delivery{}

	for rows.Next() {
		d := &models.Delivery{}
		err = rows.Scan(&d.ID, &d.CustomerID, &d.VendorID, &d.TimeOfDelivery, &d.DropLat, &d.DropLong, &d.Status)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

// GetAllByVendorID fetches all deliveries associated with a vendor id
func (m *DeliveryModel) GetAllByVendorID(vendorID int) ([]*models.Delivery, error) {
	stmt := `SELECT delivery_id, cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status
	FROM deliveries
	WHERE vendor_id = ?`

	rows, err := m.DB.Query(stmt, vendorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deliveries := []*models.Delivery{}

	for rows.Next() {
		d := &models.Delivery{}
		err = rows.Scan(&d.ID, &d.CustomerID, &d.VendorID, &d.TimeOfDelivery, &d.DropLat, &d.DropLong, &d.Status)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

// GetAllByCustomerIDStatus fetches all deliveries associated with a customer id having a particular status
func (m *DeliveryModel) GetAllByCustomerIDStatus(customerID int, status string) ([]*models.Delivery, error) {
	stmt := `SELECT delivery_id, cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status
	FROM deliveries
	WHERE cust_id = ?
	AND delivery_status = ?`

	rows, err := m.DB.Query(stmt, customerID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deliveries := []*models.Delivery{}

	for rows.Next() {
		d := &models.Delivery{}
		err = rows.Scan(&d.ID, &d.CustomerID, &d.VendorID, &d.TimeOfDelivery, &d.DropLat, &d.DropLong, &d.Status)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}

// GetAllByCustomerID fetches all deliveries associated with a customer id
func (m *DeliveryModel) GetAllByCustomerID(customerID int) ([]*models.Delivery, error) {
	stmt := `SELECT delivery_id, cust_id, vendor_id, timeofdelivery, drop_gps_lat, drop_gps_long, delivery_status
	FROM deliveries
	WHERE cust_id = ?`

	rows, err := m.DB.Query(stmt, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deliveries := []*models.Delivery{}

	for rows.Next() {
		d := &models.Delivery{}
		err = rows.Scan(&d.ID, &d.CustomerID, &d.VendorID, &d.TimeOfDelivery, &d.DropLat, &d.DropLong, &d.Status)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, d)
	}

	return deliveries, nil
}
