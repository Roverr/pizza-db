package handlers

import (
	"net/http"

	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
)

// NewStatus creates a new status structure
func NewStatus() *Status {
	return &Status{}
}

// Status describes a handler class which is responsible to handle GET /status calls
type Status struct {
}

// Handler is responsible for the request handling
func (s Status) Handler(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	// simply pong
	response.Status.Code = http.StatusOK
	response.Data = "pong"
}
