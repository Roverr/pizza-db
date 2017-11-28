package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Roverr/pizza-db/app/database/models"
)

// BulkCreatePizzas is for creating multiple pizzas in the database
func (m *Model) BulkCreatePizzas(pizzas []models.Pizza) error {
	query := `INSERT INTO pizzas (name, price) VALUES `
	for _, pizza := range pizzas {
		query = fmt.Sprintf(`%s ("%s", "%d"),`, query, pizza.Name, pizza.Price)
	}
	query = strings.TrimRight(query, ",")
	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// GetPizzas fetches all pizzas in the database
func (m *Model) GetPizzas() ([]models.Pizza, error) {
	var pizzas []models.Pizza
	query := `
  SELECT
    id, name, price
  FROM pizzas`
	if err := m.conn.Select(&pizzas, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return pizzas, nil
}

// InsertIngredientsForPizza is for inserting ingredients for a given pizza
func (m *Model) InsertIngredientsForPizza(pizzaID int64, ingredients []int64) error {
	query := `INSERT INTO pizza_ingredients (pizza_id, ingredient_id) VALUES `
	for _, ing := range ingredients {
		query = fmt.Sprintf(`%s ("%d", "%d"),`, query, pizzaID, ing)
	}
	query = strings.TrimRight(query, ",")
	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
