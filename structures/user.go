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
