package handlers

import (
	"net/http"
	"strconv"

	"github.com/Roverr/pizza-db/app/database"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// NewListIngredientsForPizza creates a new ListIngredientsForPizza structure
func NewListIngredientsForPizza(db database.Model) *ListIngredientsForPizza {
	return &ListIngredientsForPizza{db}
}

// ListIngredientsForPizza describes a handler class which is responsible to handle GET /IngredientsForPizzas calls
type ListIngredientsForPizza struct {
	db database.Model
}

// Handler is responsible for the request handling
func (lp ListIngredientsForPizza) Handler(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	pizzaID := query.Get("pizzaId")
	id, err := strconv.ParseInt(pizzaID, 10, 64)
	if err != nil {
		log.Errorln(err)
		response.Status.Code = http.StatusInternalServerError
		return
	}
	ingredients, err := lp.db.GetIngredientsForPizza(int64(id))
	if err != nil {
		log.Errorln(err)
		response.Status.Code = http.StatusInternalServerError
		return
	}
	response.Status.Code = http.StatusOK
	response.Data = ingredients
}
