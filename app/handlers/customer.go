package handlers

import (
	"net/http"

	"github.com/Roverr/pizza-db/app/database"
	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// NewListCustomer creates a new ListCustomer structure
func NewListCustomer(db database.Model) *ListCustomer {
	return &ListCustomer{db}
}

// ListCustomer describes a handler class which is responsible to handle GET /customers calls
type ListCustomer struct {
	db database.Model
}

// Handler is responsible for the request handling
func (lp ListCustomer) Handler(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	customers, err := lp.db.GetCustomers()
	if err != nil {
		log.Errorln(err)
		response.Status.Code = http.StatusInternalServerError
		return
	}
	response.Status.Code = http.StatusOK
	response.Data = customers
}
