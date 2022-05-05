package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a todo")
}

func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show details of todo %d\n", id)
}
