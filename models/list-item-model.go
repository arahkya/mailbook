package models

import "time"

type ListItemModel struct {
	Title     string    `json:"title"`
	HouseNo   string    `json:"houseNo"`
	PackageNo string    `json:"packageNo"`
	Date      time.Time `json:"createDate"`
}
