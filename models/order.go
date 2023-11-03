package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Items []Purchase `json:""`
}
