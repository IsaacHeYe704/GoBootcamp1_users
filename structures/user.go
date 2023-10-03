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
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Active   bool
	Address  Address `json:"address" validate:"required"`
}
type Address struct {
	City           string `json:"city" validate:"required"`
	Country        string `json:"country" validate:"required"`
	AddressDetails string `json:"address_details" validate:"required"`
}
