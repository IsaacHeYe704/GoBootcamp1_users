package service_test

import (
	"bootcam1_users/structures"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type StorageMock struct {
	expectedId    uuid.UUID
	expectedData  any
	expectedError error
}

func (s StorageMock) Get(id uuid.UUID) (interface{}, error) {
	if s.expectedId != id {
		return nil, errors.New("user not found")
	}
	return s.expectedData, s.expectedError

}

func (s StorageMock) GetAll() ([]interface{}, error) {
	val, _ := s.expectedData.([]interface{})
	return val, s.expectedError
}

func (s StorageMock) Create(_ uuid.UUID, _ interface{}) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (s StorageMock) Update(_ uuid.UUID, _ interface{}) (interface{}, error) {
	panic("not implemented") // TODO: Implement
}

func (s StorageMock) Delete(_ uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

var aux, _ = json.Marshal(MockUsers)
var jsonMockUser = string(aux)
var MockUsers = []interface{}{
	structures.User{
		ID: uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc2"), Name: "Isaac",
		LastName: "Herrera Yepes",
		Email:    "Isaac.herrera@globant.com",
		Active:   false,
		Address:  structures.Address{City: "Bogota", Country: "Colombia", AddressDetails: "Calle 135a ·57a 55"}},

	structures.User{
		ID:       uuid.MustParse("a56c6f0d-fe0f-49bf-9dc8-5f619c593d89"),
		Name:     "John",
		LastName: "Doe",
		Email:    "john.doe@example.com",
		Active:   false,
		Address: structures.Address{
			City:           "New York",
			Country:        "123 Main St",
			AddressDetails: "USA",
		},
	},
	structures.User{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e6"),
		Name:     "Alice",
		LastName: "Smith",
		Email:    "alice.smith@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "London",
			Country:        "United Kingdom",
			AddressDetails: "456 Oxford Street",
		},
	},

	structures.User{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e7"),
		Name:     "Bob",
		LastName: "Johnson",
		Email:    "bob.johnson@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "Los Angeles",
			Country:        "USA",
			AddressDetails: "789 Hollywood Blvd",
		},
	},
	structures.User{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e8"),
		Name:     "Emma",
		LastName: "Davis",
		Email:    "emma.davis@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "Sydney",
			Country:        "Australia",
			AddressDetails: "321 George Street",
		},
	},
}
var ExpectedUsers = []structures.User{
	{
		ID: uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc2"), Name: "Isaac",
		LastName: "Herrera Yepes",
		Email:    "Isaac.herrera@globant.com",
		Active:   false,
		Address:  structures.Address{City: "Bogota", Country: "Colombia", AddressDetails: "Calle 135a ·57a 55"}},

	{
		ID:       uuid.MustParse("a56c6f0d-fe0f-49bf-9dc8-5f619c593d89"),
		Name:     "John",
		LastName: "Doe",
		Email:    "john.doe@example.com",
		Active:   false,
		Address: structures.Address{
			City:           "New York",
			Country:        "123 Main St",
			AddressDetails: "USA",
		},
	},
	{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e6"),
		Name:     "Alice",
		LastName: "Smith",
		Email:    "alice.smith@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "London",
			Country:        "United Kingdom",
			AddressDetails: "456 Oxford Street",
		},
	},

	{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e7"),
		Name:     "Bob",
		LastName: "Johnson",
		Email:    "bob.johnson@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "Los Angeles",
			Country:        "USA",
			AddressDetails: "789 Hollywood Blvd",
		},
	},
	{
		ID:       uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e8"),
		Name:     "Emma",
		LastName: "Davis",
		Email:    "emma.davis@example.com",
		Active:   true,
		Address: structures.Address{
			City:           "Sydney",
			Country:        "Australia",
			AddressDetails: "321 George Street",
		},
	},
}
