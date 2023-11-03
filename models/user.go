package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        string    `json:""`
	FirstName string    `json:""`
	LastName  string    `json:""`
	Addresses []Address `json:""`
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Addresses: []Address{},
	}
}
