package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	gorm.Index
	StreetName string `json:"street_name"`
	Number     uint16 `json:"number"`
	City       string `json:"city"`
	ZipCode    string `json:"zip_code"`
}
