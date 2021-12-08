package main

import (
	"net/http"

	"github.com/arahkya/mailbook/handlers"
)

func main() {

	listHandler := handlers.ListHandler{}

	http.Handle("/api/list", listHandler)

	http.ListenAndServe(":18986", nil)
}
