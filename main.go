package main

import (
	"net/http"

	"github.com/arahkya/mailbook/handlers"
)

func main() {
	createHandler := handlers.CreateHandler{}
	listHandler := handlers.ListHandler{}

	http.Handle("/api/create", createHandler)
	http.Handle("/api/list", listHandler)

	http.ListenAndServe(":18986", nil)
}
