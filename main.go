package main

import (
	"bootcam1_users/structures"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	userManager := structures.NewUserManager()
	//insert user
	user, errCreating := userManager.Create(ExampleUser)
	if errCreating != nil {
		fmt.Printf("error creating user... %v", errCreating)
	}
	fmt.Printf("user created: %v", user)
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")

	//read user
	gotUser, errorReading := userManager.Get(ExampleUser.ID)
	if errorReading != nil {
		fmt.Printf("error reading user... %v", errCreating)
	}
	fmt.Printf("reading user by id %v , got user: %v ", user.ID, gotUser)
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")

	//update user
	updatedUser, errUpdating := userManager.Update(ExampleUser.ID, UpdateUser)
	if errorReading != nil {
		fmt.Printf("error reading user... %v", errUpdating)
	}
	fmt.Printf("updating user by id %v , got user: %v ", user.ID, updatedUser)
	fmt.Println("//")
	fmt.Println("//")
	fmt.Println("//")

	//read all users
	users := userManager.GetAll()
	fmt.Printf("geting all users : %v ", users)

}

var ExampleUser = structures.User{
	ID:       uuid.New(),
	Name:     "Isaac",
	LastName: "Herrera Yepes",
	Email:    "Isaac.herrera@globant.com",
	Active:   false,
	Address:  structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
var UpdateUser = structures.User{
	ID:       ExampleUser.ID,
	Name:     "updated Isaac",
	LastName: "updated Herrera Yepes",
	Email:    "updated.herrera@globant.com",
	Active:   true,
	Address:  structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
