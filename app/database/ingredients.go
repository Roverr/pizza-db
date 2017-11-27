package database

import (
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
	fmt.Println(query)
	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
