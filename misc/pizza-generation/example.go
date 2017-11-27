package main

import (
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

func generateIngredientsForDB(howMany int) []models.IngredientDB {
	ingredients := make([]models.IngredientDB, howMany)
	for i := range ingredients {
		ingredient := models.Ingredient{
			Name:       gofakeit.Name(),
			Available:  gofakeit.Bool(),
			GlutenFree: gofakeit.Bool(),
		}
		ingredients[i] = ingredient.GetDBForm()
	}
	return ingredients
}

func generate(db database.Model) error {
	var err error
	customers := generateCustomersForDB(10)
	if err = db.BulkCreateCustomer(customers); err != nil {
		return err
	}
	_, err = db.GetCustomers()
	if err != nil {
		return err
	}
	ingredients := generateIngredientsForDB(20)
	if err = db.BulkCreateIngredients(ingredients); err != nil {
		return err
	}
	return nil
}
