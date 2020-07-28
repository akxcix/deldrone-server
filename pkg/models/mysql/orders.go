package mysql

import (
	"database/sql"

	"github.com/iamadarshk/deldrone-server/pkg/models"
)

// OrderModel wraps a connection pool
type OrderModel struct {
	DB *sql.DB
}

// Insert inserts a new listing into the database using delivery id
func (m *OrderModel) Insert(deliveryID, listingID, orderQuantity, orderAmount int) error {
	stmt := `INSERT INTO orders (delivery_id, list_id, order_quantity, order_amount)
	VALUES (?, ?, ?, ?)`

	_, err := m.DB.Exec(stmt, deliveryID, listingID, orderQuantity, orderAmount)
	if err != nil {
		return err
	}

	return nil
}

// AllFromDeliveryID fetches all orders corresponding to a particular delivery ID
func (m *OrderModel) AllFromDeliveryID(deliveryID int) ([]*models.Order, error) {
	stmt := `SELECT order_id, delivery_id, list_id, order_quantity, order_amount
	FROM orders
	WHERE delivery_id = ?`

	rows, err := m.DB.Query(stmt, deliveryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*models.Order{}

	for rows.Next() {
		o := &models.Order{}
		err = rows.Scan(&o.ID, &o.DeliveryID, &o.ListingID, &o.Quant, &o.Amount)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, nil
}

// Get fetches a particular order using order_id
func (m *OrderModel) Get(orderID int) (*models.Order, error) {
	stmt := `SELECT order_id, delivery_id, list_id, order_quantity, order_amount
	FROM orders
	WHERE order_id = ?`

	row := m.DB.QueryRow(stmt, orderID)
	o := &models.Order{}
	err := row.Scan(&o.ID, &o.DeliveryID, &o.ListingID, &o.Quant, &o.Amount)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return o, nil
}
