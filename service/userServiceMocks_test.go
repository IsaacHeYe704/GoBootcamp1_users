package service_test

import (
	"bootcam1_users/structures"

	"github.com/google/uuid"
)

type StorageMock struct {
	expectedData  any
	expectedError error
}

func (s StorageMock) Get(id uuid.UUID) (interface{}, error) {
	if s.expectedError != nil {
		return nil, s.expectedError
	}
	return s.expectedData, nil

}

func (s StorageMock) GetAll() ([]interface{}, error) {
	val, _ := s.expectedData.([]interface{})
	return val, s.expectedError
}

func (s StorageMock) Create(id uuid.UUID, user interface{}) (interface{}, error) {
	if s.expectedError != nil {
		return structures.User{}, s.expectedError
	}
	return s.expectedData, nil
}

func (s StorageMock) Update(_ uuid.UUID, _ interface{}) (interface{}, error) {
	if s.expectedError != nil {
		return nil, s.expectedError
	}

	return s.expectedData, nil
}

func (s StorageMock) Delete(_ uuid.UUID) error {
	return s.expectedError
}

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
var mockCreateUserRequest = structures.UserRequest{
	Name:     "Isaac",
	LastName: "Herrera Yepes 2",
	Email:    "Isaac.herrera@globant.com",
	Active:   false,
	Address: structures.Address{
		City:           "Bogota",
		Country:        "Colombia",
		AddressDetails: "Calle 135a ·57a 55"},
}
var mockCreateUser = structures.User{
	Name:     "Isaac",
	LastName: "Herrera Yepes",
	Email:    "Isaac.herrera@globant.com",
	Active:   false,
	Address: structures.Address{
		City:           "Bogota",
		Country:        "Colombia",
		AddressDetails: "Calle 135a ·57a 55"},
}
var mockUpdatedUserRequest = structures.UserRequest{
	Name:     "Isaac updated",
	LastName: "Herrera Yepes updated",
	Email:    "updated.mail@example.com",
	Active:   false,
	Address: structures.Address{
		City:           "New York",
		Country:        "123 Main St",
		AddressDetails: "USA",
	},
}
var mockUpdatedUser = structures.User{
	Name:     "Isaac updated",
	LastName: "Herrera Yepes updated",
	Email:    "updated.mail@example.com",
	Active:   false,
	Address: structures.Address{
		City:           "New York",
		Country:        "123 Main St",
		AddressDetails: "USA",
	},
}
var mockUserJson = `{
	"id":"465f8b66-1c38-4980-b11f-aa1169f7bbc2",
	"name": "Isaac insertado",
	"lastName": "apellidaso",
	"email": "Isaac.herrera@globant.com",
	"Active": true,
	"address": {
		"city": "Bogota",
		"country": "Colombia",
		"address_details": "Calle 135a ·57a 55"
	}
}`
var mockGetUser = structures.User{
	ID:       uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc2"),
	Name:     "Isaac insertado",
	LastName: "apellidaso",
	Email:    "Isaac.herrera@globant.com",
	Active:   true,
	Address: structures.Address{
		City:           "Bogota",
		Country:        "Colombia",
		AddressDetails: "Calle 135a ·57a 55",
	},
}
