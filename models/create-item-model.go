package models

import "time"

type CreateItemModel struct {
	Title     string    `json:"title"`
	HouseNo   string    `json:"houseNo"`
	PackageNo string    `json:"packageNo"`
	Date      time.Time `json:"date"`
}
