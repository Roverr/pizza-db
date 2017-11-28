package models

// IngredientDB describes the data structure used in the database for ingredients
type IngredientDB struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Available  int8   `json:"available" db:"available"`
	GlutenFree int8   `json:"glutenFree" db:"gluten_free"`
}

// GetIngredientForm is to convert the database form into a customer facing form
func (ingd IngredientDB) GetIngredientForm() Ingredient {
	var available bool
	if ingd.Available > 0 {
		available = true
	}
	var glutenFree bool
	if ingd.GlutenFree > 0 {
		glutenFree = true
	}
	return Ingredient{
		ID:         ingd.ID,
		Name:       ingd.Name,
		Available:  available,
		GlutenFree: glutenFree,
	}
}

// Ingredient describes the data being handled by the application regarding ingredients
type Ingredient struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Available  bool   `json:"available"`
	GlutenFree bool   `json:"glutenFree"`
}

// GetDBForm is to convert a customer facing ingredient structure into the way the application
// handles it in the database
func (ing Ingredient) GetDBForm() IngredientDB {
	var available int8
	if ing.Available {
		available = 1
	}
	var glutenFree int8
	if ing.GlutenFree {
		glutenFree = 1
	}
	return IngredientDB{
		ID:         ing.ID,
		Name:       ing.Name,
		Available:  available,
		GlutenFree: glutenFree,
	}
}
