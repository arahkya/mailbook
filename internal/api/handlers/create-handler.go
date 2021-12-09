package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arahkya/mailbook/internal/api/models"
)

type CreateHandler struct{}

func (ch CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createItemModel := models.CreateItemModel{}

	jsonDecodeError := json.NewDecoder(r.Body).Decode(&createItemModel)

	if jsonDecodeError != nil {
		fmt.Println(jsonDecodeError.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
