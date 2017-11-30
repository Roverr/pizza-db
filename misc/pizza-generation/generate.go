package main

import (
	"fmt"
	"time"

	"github.com/Roverr/pizza-db/app/database"
	"github.com/Roverr/pizza-db/app/database/models"
	"github.com/brianvoe/gofakeit"
)

func getUniqueEmail(emails map[string]bool) string {
	for {
		email := gofakeit.Email()
		ok := emails[email]
		if !ok {
			emails[email] = true
			return email
		}
	}
}

func generateCustomersForDB(howMany int) []models.CustomerDB {
	customers := make([]models.CustomerDB, howMany)
	emailChecker := map[string]bool{}
	for i := range customers {
		customers[i].Email = getUniqueEmail(emailChecker)
		customers[i].Name = gofakeit.Name()
		customers[i].Password = gofakeit.Password(true, true, true, true, true, 32)
	}
	return customers
}

func generateIngredientsForDB(howMany int, allAvailable bool) []models.IngredientDB {
	ingredients := make([]models.IngredientDB, howMany)
	for i := range ingredients {
		available := gofakeit.Bool()
		if allAvailable {
			available = true
		}
		ingredient := models.Ingredient{
			Name:       gofakeit.BeerName() + gofakeit.BuzzWord(),
			Available:  available,
			GlutenFree: gofakeit.Bool(),
		}
		ingredients[i] = ingredient.GetDBForm()
	}
	return ingredients
}

func generatePizzasForDB(howMany int) []models.Pizza {
	pizzas := make([]models.Pizza, howMany)
	for i := range pizzas {
		pizzas[i].Name = fmt.Sprintf("%s %s", gofakeit.City(), gofakeit.JobLevel())
		pizzas[i].Price = int64(gofakeit.Number(5, 70))
	}
	return pizzas
}

func twoRandomNumbersInRange(a int, b int) (int, int) {
	for {
		i := gofakeit.Number(a, b)
		y := gofakeit.Number(a, b)
		if i != y {
			return i, y
		}
	}
}

// Generate is the main function for generating random data and inserting it into the database
func generate(db database.Model) error {
	var err error

	// Generate customers
	customers := generateCustomersForDB(10)
	if err = db.BulkCreateCustomer(customers); err != nil {
		return err
	}
	customers, err = db.GetCustomers()
	if err != nil {
		return err
	}

	// Generate Pizzas
	pizzas := generatePizzasForDB(8)
	if err = db.BulkCreatePizzas(pizzas); err != nil {
		return err
	}
	pizzas, err = db.GetPizzas()
	if err != nil {
		return err
	}

	// Generate ingredients for the pizzas
	for _, pizza := range pizzas {
		// Generate ingredients
		ingredientsDB := generateIngredientsForDB(5, gofakeit.Bool())
		if err = db.BulkCreateIngredients(ingredientsDB); err != nil {
			return err
		}
		ingredients, err := db.GetLastIngredients(5)
		if err != nil {
			return err
		}
		indexes := []int64{}
		for _, ing := range ingredients {
			indexes = append(indexes, ing.ID)
		}
		err = db.InsertIngredientsForPizza(pizza.ID, indexes)
		if err != nil {
			return err
		}
	}

	// Generate orders
	for _, cust := range customers {
		started := gofakeit.DateRange(time.Now().Truncate(time.Hour*368), time.Now())
		completed := started.Add(time.Hour)
		order := models.Order{
			Price:       int64(gofakeit.Number(30, 50)),
			Address:     gofakeit.Address().Address,
			StartedAt:   started,
			CompletedAt: &completed,
			CustomerID:  cust.ID,
		}
		first, second := twoRandomNumbersInRange(0, len(pizzas)-1)
		details := []models.OrderDetail{
			models.OrderDetail{
				PizzaType: pizzas[first].ID,
				HowMany:   int8(gofakeit.Number(1, 3)),
			},
			models.OrderDetail{
				PizzaType: pizzas[second].ID,
				HowMany:   int8(gofakeit.Number(1, 3)),
			},
		}
		err := db.InsertOrder(order, details)
		if err != nil {
			return err
		}
	}
	return nil
}
