package main

import (
	"net/http"
	"path/filepath"

	"github.com/arahkya/mailbook/internal/web/handlers"
)

func main() {
	scriptFolderPath, _ := filepath.Abs("../../assets/web/wwwroot")

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(scriptFolderPath))))
	http.Handle("/", handlers.IndexHandler{})

	http.ListenAndServe(":28986", nil)
}
