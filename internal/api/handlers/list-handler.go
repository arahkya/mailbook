package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/arahkya/mailbook/internal/api/models"
)

type ListHandler struct{}

func (lh ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	listItems := []models.ListItemModel{}

	listItems = append(listItems,
		models.ListItemModel{
			Title:     "Package 0001",
			HouseNo:   "89/86",
			PackageNo: "PKG0001",
			Date:      time.Now(),
		},
		models.ListItemModel{
			Title:     "Package 0002",
			HouseNo:   "89/86",
			PackageNo: "PKG0002",
			Date:      time.Now(),
		},
		models.ListItemModel{
			Title:     "Package 0003",
			HouseNo:   "89/86",
			PackageNo: "PKG0003",
			Date:      time.Now(),
		},
		models.ListItemModel{
			Title:     "Package 0004",
			HouseNo:   "89/86",
			PackageNo: "PKG0004",
			Date:      time.Now(),
		},
	)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(&listItems)
}
