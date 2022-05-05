package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// health check endpoint
	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthCheckHandler)

	// todos endpoint
	router.HandlerFunc(http.MethodPost, "/api/v1/todos", app.createTodoHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/todos/:id", app.showTodoHandler)

	return router
}
