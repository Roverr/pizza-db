package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Roverr/pizza-db/app/database/models"
)

// BulkCreateIngredients is for creating multiple ingredients
func (m *Model) BulkCreateIngredients(ingredients []models.IngredientDB) error {
	query := `INSERT INTO ingredients (name, available, gluten_free) VALUES `
	for _, ing := range ingredients {
		query = fmt.Sprintf(`%s ("%s", "%d", "%d"),`, query, ing.Name, ing.Available, ing.GlutenFree)
	}
	query = strings.TrimRight(query, ",")
	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// GetIngredients is to query all ingredients found in the database
func (m *Model) GetIngredients() ([]models.Ingredient, error) {
	var ingredientsDB []models.IngredientDB
	query := `
  SELECT
    id, name, available, gluten_free
  FROM ingredients`
	if err := m.conn.Select(&ingredientsDB, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var ingredients []models.Ingredient
	for _, ing := range ingredientsDB {
		ingredients = append(ingredients, ing.GetIngredientForm())
	}
	return ingredients, nil
}
