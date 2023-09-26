package structures

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	LastName string
	Email    string
	Active   bool
	Address  Address
}
type UserRequest struct {
	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Email    string `validate:"required,email"`
	Active   bool
	Address  Address `validate:"required"`
}
type Address struct {
	City           string `validate:"required"`
	Country        string `validate:"required"`
	AddressDetails string `validate:"required"`
}
