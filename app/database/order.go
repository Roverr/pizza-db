package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Roverr/pizza-db/app/database/models"
	log "github.com/sirupsen/logrus"
)

// Order describes an order made by a customer
type Order struct {
	ID          int64      `json:"id" db:"id"`
	CustomerID  int64      `json:"customerId" db:"customer_id"`
	StartedAt   time.Time  `json:"startedAt" db:"started_at"`
	CompletedAt *time.Time `json:"completedAt" db:"completed_at"`
	Price       int64      `json:"price" db:"price"`
	Address     string     `json:"address" db:"address"`
}

// InsertOrder is for inserting orders made by customers
func (m *Model) InsertOrder(order models.Order, details []models.OrderDetail) error {
	tx, err := m.conn.Begin()
	if err != nil {
		return err
	}
	var result sql.Result
	result, err = tx.Exec(`
    INSERT INTO orders (customer_id, started_at, completed_at, price, address)
    VALUES (?, ?, ?, ?, ?)`,
		order.CustomerID,
		order.StartedAt,
		order.CompletedAt,
		order.Price,
		order.StartedAt,
	)
	if err != nil {
		log.Error(err)
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		log.Error(err)
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}
	query := `INSERT INTO pizzas_for_orders (order_id, pizza_id, number_of_pizzas) VALUES `
	for _, detail := range details {
		query = fmt.Sprintf(`%s ("%d", "%d", "%d"),`, query, id, detail.PizzaType, detail.HowMany)
	}
	query = strings.TrimRight(query, ",")
	_, err = tx.Exec(query)
	if err != nil {
		log.Error(err)
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}
	return tx.Commit()
}
