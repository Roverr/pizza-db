package main

import (
	"net/http"

	"github.com/Roverr/pizza-db/app/config"
	"github.com/Roverr/pizza-db/app/database"
	"github.com/Roverr/pizza-db/app/handlers"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/TobiEiss/goMiddlewareChain/templates"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	var db *database.Model
	db, err = database.New(&config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	status := handlers.NewStatus()
	router.GET("/api/status", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, status.Handler))

	listPizza := handlers.NewListPizza(*db)
	router.GET("/api/pizzas", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, listPizza.Handler))

	listCustomer := handlers.NewListCustomer(*db)
	router.GET("/api/customers", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, listCustomer.Handler))

	listOrders := handlers.NewListOrder(*db)
	router.GET("/api/orders", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, listOrders.Handler))

	listIngredientsForPizzas := handlers.NewListIngredientsForPizza(*db)
	router.GET("/api/ingredients", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, listIngredientsForPizzas.Handler))

	router.ServeFiles("/ui/*filepath", http.Dir("./public"))
	handler := cors.Default().Handler(router)
	log.Infof("Application is listening on: %s", config.ListenAddress)
	log.Fatal(http.ListenAndServe(config.ListenAddress, handler))
}
