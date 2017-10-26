package app

import (
	"net/http"

	"github.com/TobiEiss/goMiddlewareChain"
	"github.com/julienschmidt/httprouter"
)

// Status is a Handler for GET /status endpoint, returns a ping
func Status(response *goMiddlewareChain.Response, request *http.Request, params httprouter.Params) {
	response.Status.Code = http.StatusOK
	response.Data = "pong"
}
