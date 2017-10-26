package main

import (
	"log"
	"net/http"

	"github.com/Roverr/pizza-db/app/config"
	"github.com/Roverr/pizza-db/app/database"
	"github.com/Roverr/pizza-db/app/handlers"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/TobiEiss/goMiddlewareChain/templates"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.New(&config.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	statusHandler := handlers.NewStatus()
	router.GET("/api/status", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, statusHandler.Handler))
	router.ServeFiles("/ui/*filepath", http.Dir("./public"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
