package structures

import "github.com/google/uuid"

var DefaultUsers = []User{
	User{
		"465f8b66-1c38-4980-b11f-aa1169f7bbc2", "Isaac",
		"Herrera Yepes",
		"Isaac.herrera@globant.com",
		false,
		Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}},
	User{
		"a56c6f0d-fe0f-49bf-9dc8-5f619c593d89",
		"John",
		"Doe",
		"john.doe@example.com",
		false,
		Address{
			"New York",
			"USA",
			"123 Main St",
		},
	},
	User{
		"c20ba804-122f-4063-bb09-6cbfba6a28e6",
		"Alice",
		"Smith",
		"alice.smith@example.com",
		true,
		Address{
			"London",
			"United Kingdom",
			"456 Oxford Street",
		},
	},

	User{
		uuid.NewString(),
		"Bob",
		"Johnson",
		"bob.johnson@example.com",
		true,
		Address{
			"Los Angeles",
			"USA",
			"789 Hollywood Blvd",
		},
	},
	User{
		uuid.NewString(),
		"Emma",
		"Davis",
		"emma.davis@example.com",
		true,
		Address{
			"Sydney",
			"Australia",
			"321 George Street",
		},
	},
}
