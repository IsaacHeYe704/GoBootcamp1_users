package structures_test

import (
	"bootcam1_users/structures"
	"testing"

	"github.com/google/uuid"
)

func TestRead(t *testing.T) {
	t.Run("test get", func(t *testing.T) {
		testUuid := structures.DefaultUsers[2].ID
		expectedUser := structures.DefaultUsers[2]
		userManager, _ := structures.NewUserManager()
		user, _ := userManager.Read(testUuid)

		if user != expectedUser {
			t.Errorf("expected %v and got %v", expectedUser, &user)
		}
	})
	t.Run("test get not existing uudi", func(t *testing.T) {
		userManager, _ := structures.NewUserManager()
		_, err := userManager.Read("this_uuid_does_not_exist")
		expectedError := structures.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
}
func TestCreate(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		testUser := structures.User{
			"465f8b66-1c38-4980-b11f-aa1169f7bbc3", "Inserted",
			"Herrera Yepes",
			"Isaac.herreraInserted@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager, _ := structures.NewUserManager()
		gotUser, _ := userManager.Create(testUser)

		if testUser != gotUser {
			t.Errorf("expected %v and got %v", testUser, gotUser)
		}
	})
	t.Run("create an user with existing uuid should return an error", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		duplicatedUser := structures.DefaultUsers[1]
		userManager, _ := structures.NewUserManager()
		_, err := userManager.Create(duplicatedUser)
		expectErr := structures.Error_UuidAlreadyExists
		if err != expectErr {
			t.Errorf("expected  error: %v but, got %v", expectErr, err)
		}
	})
}

var updatedUser = structures.User{
	uuid.NewString(),
	"updated name",
	"updated lastname",
	"updated@example.com",
	false,
	structures.Address{
		"New York",
		"USA",
		"123 Main St",
	},
}

func TestUpdate(t *testing.T) {
	t.Run("update user", func(t *testing.T) {

		updatedUser := structures.User{
			"465f8b66-1c38-4980-b11f-aa1169f7bbc2", "updated",
			"updated Herrera Yepes",
			"Isaac.herreraUpdated@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager, _ := structures.NewUserManager()
		gotUser, _ := userManager.Update(updatedUser.ID, updatedUser)

		if updatedUser != gotUser {
			t.Errorf("expected %v and got %v", updatedUser, gotUser)
		}
	})
	t.Run("should not update a not existing uuid", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		doesntExist := structures.User{
			"this_id_does_not_exist", "NA",
			"NA last name",
			"NA@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager, _ := structures.NewUserManager()
		_, err := userManager.Update(doesntExist.ID, doesntExist)
		expectErr := structures.Error_UserNotFound
		if err != expectErr {
			t.Errorf("expected  error: %v but, got %v", expectErr, err)
		}
	})
}
func TestDelete(t *testing.T) {
	t.Run("delete user", func(t *testing.T) {
		testUser := structures.DefaultUsers[0]
		userManager, _ := structures.NewUserManager()
		userManager.Delete(testUser.ID)
		_, err := userManager.Read(testUser.ID)
		expectedError := structures.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
	t.Run("should not delete a not existing uuid", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		testUser := "this_uuid_does_not_exist"
		userManager, _ := structures.NewUserManager()
		err := userManager.Delete(testUser)
		expectedError := structures.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
}
