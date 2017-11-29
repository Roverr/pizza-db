package handlers

import (
	"net/http"

	"github.com/Roverr/pizza-db/app/database"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
)

// NewListPizza creates a new ListPizza structure
func NewListPizza(db database.Model) *ListPizza {
	return &ListPizza{db}
}

// ListPizza describes a handler class which is responsible to handle GET /pizza calls
type ListPizza struct {
	db database.Model
}

// Handler is responsible for the request handling
func (lp ListPizza) Handler(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	pizzas, err := lp.db.GetExtendedPizzas()
	if err != nil {
		response.Status.Code = http.StatusInternalServerError
		return
	}
	response.Status.Code = http.StatusOK
	response.Data = pizzas
}
