package db

import (
	"bootcam1_users/structures"

	"github.com/google/uuid"
)

var DefaultUsers = []structures.User{
	{
		uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc2"), "Isaac",
		"Herrera Yepes",
		"Isaac.herrera@globant.com",
		false,
		structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}},

	{
		uuid.MustParse("a56c6f0d-fe0f-49bf-9dc8-5f619c593d89"),
		"John",
		"Doe",
		"john.doe@example.com",
		false,
		structures.Address{
			"New York",
			"123 Main St",
			"USA",
		},
	},
	{
		uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e6"),
		"Alice",
		"Smith",
		"alice.smith@example.com",
		true,
		structures.Address{
			"London",
			"United Kingdom",
			"456 Oxford Street",
		},
	},

	{
		uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e7"),
		"Bob",
		"Johnson",
		"bob.johnson@example.com",
		true,
		structures.Address{
			"Los Angeles",
			"USA",
			"789 Hollywood Blvd",
		},
	},
	{
		uuid.MustParse("c20ba804-122f-4063-bb09-6cbfba6a28e8"),
		"Emma",
		"Davis",
		"emma.davis@example.com",
		true,
		structures.Address{
			"Sydney",
			"Australia",
			"321 George Street",
		},
	},
}
