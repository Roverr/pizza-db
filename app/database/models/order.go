package models

import "time"

// Order describes an order made by a customer
type Order struct {
	ID          int64      `json:"id" db:"id"`
	CustomerID  int64      `json:"customerId" db:"customer_id"`
	StartedAt   time.Time  `json:"startedAt" db:"started_at"`
	CompletedAt *time.Time `json:"completedAt" db:"completed_at"`
	Price       int64      `json:"price" db:"price"`
	Address     string     `json:"address" db:"address"`
}

// OrderDetail describes the number of pizzas in a given order
type OrderDetail struct {
	HowMany   int8  // indicates how many pizza is being ordered
	PizzaType int64 // indicates the ID of the pizza being ordered
}
