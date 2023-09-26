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
	Name     string
	LastName string
	Email    string
	Active   bool
	Address  Address
}
type Address struct {
	City           string
	Country        string
	AddressDetails string
}
