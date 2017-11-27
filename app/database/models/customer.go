package models

// Customer describes the raw structure how the application handles customers
type Customer struct {
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// CustomerDB describes the raw data structure stored in the database
type CustomerDB struct {
	Customer
	ID       int64  `json:"id" db:"id"`
	Password string `json:"-" db:"password"`
}
