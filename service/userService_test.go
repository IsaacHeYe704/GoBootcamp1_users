package service_test

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/db"
	"bootcam1_users/service"
	"bootcam1_users/structures"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

//fill the exted array with users from the mock

func TestGetAllService(t *testing.T) {
	testTable := []struct {
		name           string
		storageMock    db.Storage
		expectedResult []structures.User
		expectedError  error
	}{
		{
			name: "Should get all Users",
			storageMock: StorageMock{
				expectedData:  MockUsers,
				expectedError: nil,
			},
			expectedResult: ExpectedUsers,
			expectedError:  nil,
		},
		{
			name: "Should get an error if  storage conection could not be stablish",
			storageMock: StorageMock{
				expectedData:  make([]interface{}, 0),
				expectedError: errors.New("connection refused"),
			},
			expectedResult: nil,
			expectedError: custom_errors.ServiceError{
				Code:        "ConectionError",
				Description: "connection refused",
			},
		},
		{
			name: "Should get an empty array if there are no users",
			storageMock: StorageMock{
				expectedData:  make([]interface{}, 0),
				expectedError: nil,
			},
			expectedResult: make([]structures.User, 0),
			expectedError:  nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//
			service := service.NewUserService(test.storageMock)
			got, err := service.GetAll()

			if !errors.Is(err, test.expectedError) {
				fmt.Println(err)
				t.Fatalf("Expected error: \n %v , \n  but got error: \n %v", test.expectedError, err)
			}
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Expected \n %v, \n  but got \n %v", test.expectedResult, got)
			}
		})
	}

}

func TestGetService(t *testing.T) {
	testTable := []struct {
		name           string
		storageMock    db.Storage
		expectedResult structures.User
		expectedError  error
	}{
		{
			name: "Should get an user by id",
			storageMock: StorageMock{
				expectedId:    ExpectedUsers[0].ID,
				expectedData:  MockUsers[0],
				expectedError: nil,
			},
			expectedResult: ExpectedUsers[0],
			expectedError:  nil,
		},
		{
			name: "Should get an user by id",
			storageMock: StorageMock{
				expectedId:    ExpectedUsers[1].ID,
				expectedData:  MockUsers[1],
				expectedError: nil,
			},
			expectedResult: ExpectedUsers[1],
			expectedError:  nil,
		},
		{
			name: "Should get an user not found error",
			storageMock: StorageMock{
				expectedId:    ExpectedUsers[1].ID,
				expectedData:  nil,
				expectedError: errors.New("user not found"),
			},
			expectedResult: structures.User{},
			expectedError: custom_errors.ServiceError{
				Code:        "NotFound",
				Description: "user not found",
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//
			service := service.NewUserService(test.storageMock)
			got, err := service.Get(test.expectedResult.ID)

			if !errors.Is(err, test.expectedError) {
				t.Fatalf("Expected error: \n %v , \n  but got error: \n %v", test.expectedError, err)
			}
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Expected \n %v, \n  but got \n %v", test.expectedResult, got)
			}
		})
	}

}
