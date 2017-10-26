package main

import (
	"log"
	"net/http"

	"github.com/Roverr/pizza-db/app"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/TobiEiss/goMiddlewareChain/templates"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/", http.Dir("/public"))
	router.GET("/status", goMiddlewareChain.RequestChainHandler(templates.JSONResponseHandler, app.Status))

	log.Fatal(http.ListenAndServe(":8080", router))
}
