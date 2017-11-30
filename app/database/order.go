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
		order.Address,
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

// GetOrders lists all the orders made by useres
func (m *Model) GetOrders() ([]models.ExtendedOrder, error) {
	var dbData []struct {
		OrderPrice    int64      `db:"order_price"`
		OrderAddress  string     `db:"order_address"`
		OrderID       int64      `db:"order_id"`
		CustomerName  string     `db:"customer_name"`
		CustomerID    int64      `db:"customer_id"`
		PizzaName     string     `db:"pizza_name"`
		NumberOfPizza int8       `db:"number_of_pizzas"`
		StartedAt     time.Time  `db:"started_at"`
		CompletedAt   *time.Time `db:"completed_at"`
	}
	query := `
  SELECT
    orders.price as order_price,
    orders.address as order_address,
    orders.id as order_id,
    orders.started_at as started_at,
    orders.completed_at as completed_at,
    customers.name as customer_name,
    customers.id as customer_id,
    pizzas.name as pizza_name,
    pizzas_for_orders.number_of_pizzas as number_of_pizzas
  FROM customers
  JOIN orders on orders.customer_id = customers.id
  JOIN pizzas_for_orders on pizzas_for_orders.order_id = orders.id
  JOIN pizzas on pizzas_for_orders.pizza_id = pizzas.id;`
	err := m.conn.Select(&dbData, query)
	if err != nil {
		return nil, err
	}

	orders := map[int64]models.ExtendedOrder{}
	pizzaNumbers := map[int64][]models.PizzaNumber{}
	for _, data := range dbData {
		orderID := data.OrderID
		pNumber := models.PizzaNumber{
			Name:   data.PizzaName,
			Number: data.NumberOfPizza,
		}
		_, ok := orders[orderID]
		if !ok {
			orders[orderID] = models.ExtendedOrder{
				ID: orderID,
				Customer: models.IdentifiedCustomer{
					ID:   data.CustomerID,
					Name: data.CustomerName,
				},
				StartedAt:   data.StartedAt,
				CompletedAt: data.CompletedAt,
				Address:     data.OrderAddress,
				Price:       data.OrderPrice,
				Pizzas:      []models.PizzaNumber{},
			}
		}
		element, ok := pizzaNumbers[orderID]
		if !ok {
			pizzaNumbers[orderID] = []models.PizzaNumber{pNumber}
		}
		pizzaNumbers[orderID] = append(element, pNumber)
	}
	for id := range orders {
		element := orders[id]
		element.Pizzas = pizzaNumbers[id]
		orders[id] = element
	}
	var arrayConv []models.ExtendedOrder
	for _, order := range orders {
		arrayConv = append(arrayConv, order)
	}
	return arrayConv, nil
}
