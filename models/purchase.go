package models

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	What    Item `json:""`
	HowMany int  `json:""`
}
