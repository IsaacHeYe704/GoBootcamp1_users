package db_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"

	"bootcam1_users/custom_errors"
	"bootcam1_users/db"
	"bootcam1_users/structures"
)

func TestLocalGet(t *testing.T) {
	t.Run("test get", func(t *testing.T) {
		testUuid := db.DefaultUsers[2].ID
		expectedUser := db.DefaultUsers[2]
		userManager := db.NewLocalStorage()
		user, _ := userManager.Get(testUuid)

		if user != expectedUser {
			t.Errorf("expected %v and got %v", expectedUser, &user)
		}
	})
	t.Run("test get not existing uudi", func(t *testing.T) {
		userManager := db.NewLocalStorage()
		_, err := userManager.Get(uuid.UUID{})
		expectedError := custom_errors.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
}
func TestLocalGetAll(t *testing.T) {
	t.Run("test get", func(t *testing.T) {

		userManager := db.NewLocalStorage()
		users, _ := userManager.GetAll()
		expectedUsers := make(map[uuid.UUID]structures.User)
		for _, user := range db.DefaultUsers {
			expectedUsers[user.ID] = user
		}
		if !reflect.DeepEqual(expectedUsers, users) {
			t.Errorf("expected:  %v ,and got: %v", expectedUsers, users)
		}
	})
	t.Run("test get not existing uudi", func(t *testing.T) {
		userManager := db.NewLocalStorage()
		_, err := userManager.Get(uuid.UUID{})
		expectedError := custom_errors.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
}
func TestLocalCreate(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		testUser := structures.User{
			uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc3"), "Inserted",
			"Herrera Yepes",
			"Isaac.herreraInserted@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager := db.NewLocalStorage()
		gotUser, _ := userManager.Create(testUser)

		if testUser != gotUser {
			t.Errorf("expected %v and got %v", testUser, gotUser)
		}
	})
	t.Run("create an user with existing uuid should return an error", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		duplicatedUser := db.DefaultUsers[1]
		userManager := db.NewLocalStorage()
		_, err := userManager.Create(duplicatedUser)
		expectErr := custom_errors.Error_UuidAlreadyExists
		if err != expectErr {
			t.Errorf("expected  error: %v but, got %v", expectErr, err)
		}
	})
}

var updatedUser = structures.User{
	uuid.MustParse(uuid.NewString()),
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

func TestLocalUpdate(t *testing.T) {
	t.Run("update user", func(t *testing.T) {

		updatedUser := structures.User{
			uuid.MustParse("465f8b66-1c38-4980-b11f-aa1169f7bbc2"), "updated",
			"updated Herrera Yepes",
			"Isaac.herreraUpdated@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager := db.NewLocalStorage()
		gotUser, _ := userManager.Update(updatedUser.ID, updatedUser)

		if updatedUser != gotUser {
			t.Errorf("expected %v and got %v", updatedUser, gotUser)
		}
	})
	t.Run("should not update a not existing uuid", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		doesntExist := structures.User{
			uuid.UUID{}, "NA",
			"NA last name",
			"NA@globant.com",
			false,
			structures.Address{"Bogota", "Colombia", "Calle 135a ·57a 55"}}
		userManager := db.NewLocalStorage()
		_, err := userManager.Update(doesntExist.ID, doesntExist)
		expectErr := custom_errors.Error_UserNotFound
		if err != expectErr {
			t.Errorf("expected  error: %v but, got %v", expectErr, err)
		}
	})
}
func TestLocalDelete(t *testing.T) {
	t.Run("delete user", func(t *testing.T) {
		testUser := db.DefaultUsers[0]
		userManager := db.NewLocalStorage()
		userManager.Delete(testUser.ID)
		_, err := userManager.Get(testUser.ID)
		expectedError := custom_errors.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
	t.Run("should not delete a not existing uuid", func(t *testing.T) {
		//an user with this id is already inserted in the userManagment creatin
		testUuid := uuid.UUID{}
		userManager := db.NewLocalStorage()
		err := userManager.Delete(testUuid)
		expectedError := custom_errors.Error_UserNotFound
		if expectedError != err {
			t.Errorf("expected %v and got %v", expectedError, err)
		}
	})
}
