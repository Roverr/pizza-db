package main

import (
	"github.com/Roverr/pizza-db/app/database"
	log "github.com/sirupsen/logrus"
)

func clean(db database.Model) error {
	custQuery := `DELETE FROM customers`
	ingQuery := `DELETE FROM ingredients`
	pizzQuery := `DELETE FROM pizzas`
	ordQuery := `DELETE FROM orders`

	conn := db.GetConnection()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.Exec(custQuery); err != nil {
		log.Error(err)
		return tx.Rollback()
	}
	if _, err = tx.Exec(ingQuery); err != nil {
		log.Error(err)
		return tx.Rollback()
	}
	if _, err = tx.Exec(pizzQuery); err != nil {
		log.Error(err)
		return tx.Rollback()
	}
	if _, err = tx.Exec(ordQuery); err != nil {
		log.Error(err)
		return tx.Rollback()
	}
	return tx.Commit()
}
