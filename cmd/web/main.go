package main

import (
	"net/http"

	"github.com/arahkya/mailbook/internal/web/handlers"
)

func main() {
	http.Handle("/", handlers.IndexHandler{})

	http.ListenAndServe(":28986", nil)
}
