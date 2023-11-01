package models

type User struct {
	Id        string
	FirstName string
	LastName  string
	Addresses []Address
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Addresses: []Address{},
	}
}
