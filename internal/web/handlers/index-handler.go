package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type IndexHandler struct{}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	templateFolderPath, _ := filepath.Abs("../../assets/web")
	templateFilePath := filepath.Join(templateFolderPath, "*.gohtml")
	templateFiles, parseGlogError := template.ParseGlob(templateFilePath)

	if parseGlogError != nil {
		log.Fatalln(parseGlogError.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	templateFiles.Lookup("index.gohtml").Execute(w, nil)
}
