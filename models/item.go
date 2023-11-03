package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id           string `json:""`
	Name         string `json:""`
	Description  string `json:""`
	isVeggie     bool   `json:""`
	isGlutenFree bool   `json:""`
}
