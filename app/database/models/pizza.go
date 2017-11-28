package models

// Pizza describes a pizza's structure
type Pizza struct {
	ID    int64  `json:"id" db:"id"`
	Price int64  `json:"price" db:"price"`
	Name  string `json:"name" db:"name"`
}

// PizzaIngredient describes an ingredient which is for a given pizza
type PizzaIngredient struct {
	PizzaID      int64 `json:"pizzaId" db:"pizza_id"`
	IngredientID int64 `json:"ingredientId" db:"ingredient_id"`
}
