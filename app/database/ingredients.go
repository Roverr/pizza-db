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

// GetIngredientsForPizza is for getting ingredients for a given pizza
func (m *Model) GetIngredientsForPizza(pizzaID int64) ([]models.Ingredient, error) {
	var ingredientsDB []models.IngredientDB
	query := `
  SELECT
    ingredients.id,
    ingredients.name,
    ingredients.available,
    ingredients.gluten_free
  FROM ingredients
  JOIN pizza_ingredients on pizza_ingredients.ingredient_id = ingredients.id
  WHERE pizza_ingredients.pizza_id=?;`
	if err := m.conn.Select(&ingredientsDB, query, pizzaID); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var ingredients []models.Ingredient
	for _, ing := range ingredientsDB {
		ingredients = append(ingredients, ing.GetIngredientForm())
	}
	return ingredients, nil
}

// GetLastIngredients is for getting the last given number of ingredients from db
func (m *Model) GetLastIngredients(limit int) ([]models.Ingredient, error) {
	var ingredientsDB []models.IngredientDB
	query := fmt.Sprintf(`
  SELECT
    id, name, available, gluten_free
  FROM ingredients
  ORDER by id desc
  LIMIT %d`, limit)
	if err := m.conn.Select(&ingredientsDB, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var ingredients []models.Ingredient
	for _, ing := range ingredientsDB {
		ingredients = append(ingredients, ing.GetIngredientForm())
	}
	return ingredients, nil
}
