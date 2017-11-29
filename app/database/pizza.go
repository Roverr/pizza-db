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

// GetExtendedPizzas gives back a structure which describes the necessary ingredients for every pizza
func (m *Model) GetExtendedPizzas() ([]models.ExtendedPizza, error) {
	var comps []struct {
		IngredientID   int64  `db:"ing_id"`
		IngredientName string `db:"ing_name"`
		Available      bool   `db:"available"`
		GlutenFree     bool   `db:"gluten_free"`
		PizzaName      string `db:"pizza_name"`
		PizzaPrice     int64  `db:"pizza_price"`
		PizzaID        int64  `db:"pizza_id"`
	}
	query := `
  SELECT
    ingredients.id as ing_id,
    ingredients.name as ing_name,
    ingredients.available as available,
    ingredients.gluten_free as gluten_free,
    ingredients.id as ing_id,
    pizzas.name as pizza_name,
    pizzas.price as pizza_price,
    pizzas.id as pizza_id
  FROM ingredients
  JOIN pizza_ingredients on pizza_ingredients.ingredient_id = ingredients.id
  JOIN pizzas on pizzas.id = pizza_ingredients.pizza_id;`
	if err := m.conn.Select(&comps, query); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	shortcut := map[int64][]int{}
	var extendeds []models.ExtendedPizza
	for i, comp := range comps {
		_, ok := shortcut[comp.PizzaID]
		if !ok {
			extended := models.ExtendedPizza{
				Pizza: models.Pizza{
					Name:  comp.PizzaName,
					ID:    comp.PizzaID,
					Price: comp.PizzaPrice,
				},
				Ingrients: []models.Ingredient{},
			}
			extendeds = append(extendeds, extended)
			shortcut[comp.PizzaID] = []int{i}
			continue
		}
		shortcut[comp.PizzaID] = append(shortcut[comp.PizzaID], i)
	}

	for i, extended := range extendeds {
		ids := shortcut[extended.Pizza.ID]
		for _, id := range ids {
			ingredient := models.Ingredient{
				ID:         comps[id].IngredientID,
				Name:       comps[id].IngredientName,
				GlutenFree: comps[id].GlutenFree,
				Available:  comps[id].Available,
			}
			extendeds[i].Ingrients = append(extendeds[i].Ingrients, ingredient)
		}
	}
	return extendeds, nil
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
