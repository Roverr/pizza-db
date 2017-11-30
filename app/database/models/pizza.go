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

// ExtendedPizza describes a pizza which has been extended with ingredients informations
type ExtendedPizza struct {
	Pizza
	Ingrients []Ingredient `json:"ingredients"`
}

// PizzaNumber describes a pizza's name and the number of the pizzas in the order from that type
type PizzaNumber struct {
	Name   string `json:"name"`
	Number int8   `json:"number"`
}
