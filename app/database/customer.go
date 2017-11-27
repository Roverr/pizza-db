package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Roverr/pizza-db/app/database/models"
	"github.com/davecgh/go-spew/spew"
)

// BulkCreateCustomer is for creating multiple customers
func (m *Model) BulkCreateCustomer(customers []models.CustomerDB) error {
	query := `INSERT INTO customers (email, name, password) VALUES `
	for _, cust := range customers {
		query = fmt.Sprintf(`%s ("%s", "%s", "%s"),`, query, cust.Email, cust.Name, cust.Password)
	}
	query = strings.TrimRight(query, ",")
	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// GetCustomers is for getting all customers from the database
func (m *Model) GetCustomers() ([]models.CustomerDB, error) {
	var customers []models.CustomerDB
	query := `
  SELECT
    id, name, email, password
  FROM customers`
	if err := m.conn.Select(&customers, query); err != nil && err != sql.ErrNoRows {
		spew.Dump(err)
		return nil, err
	}
	return customers, nil
}
