package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes is used to register the handlers
func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	// Convert the notFoundResponse helper to a http.Handler.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// Convert the methodNotAllowedResponse to customer error handler for 405 responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/api/message", app.createMessageHandler)
	router.HandlerFunc(http.MethodGet, "/api/hash/:hash", app.showMessageHandler)
	return router
}
