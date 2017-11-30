package handlers

import (
	"net/http"

	"github.com/Roverr/pizza-db/app/database"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// NewListOrder creates a new ListOrder structure
func NewListOrder(db database.Model) *ListOrder {
	return &ListOrder{db}
}

// ListOrder describes a handler class which is responsible to handle GET /orders calls
type ListOrder struct {
	db database.Model
}

// Handler is responsible for the request handling
func (lp ListOrder) Handler(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	orders, err := lp.db.GetOrders()
	if err != nil {
		log.Errorln(err)
		response.Status.Code = http.StatusInternalServerError
		return
	}
	response.Status.Code = http.StatusOK
	response.Data = orders
}
