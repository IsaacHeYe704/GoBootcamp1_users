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

	"github.com/google/uuid"
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
				getAllMock: func() ([]interface{}, error) {
					return MockUsers, nil
				},
			},
			expectedResult: ExpectedUsers,
			expectedError:  nil,
		},
		{
			name: "Should get an error if  storage conection could not be stablish",
			storageMock: StorageMock{
				getAllMock: func() ([]interface{}, error) {
					return nil, errors.New("connection refused")
				},
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
				getAllMock: func() ([]interface{}, error) {
					return make([]interface{}, 0), nil
				},
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
				t.Fatalf("Expected error: \n %d , \n  but got error: \n %d", test.expectedError, err)
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
				getByIdMock: func() (interface{}, error) {
					return MockUsers[0], nil
				},
			},
			expectedResult: ExpectedUsers[0],
			expectedError:  nil,
		},
		{
			name: "Should get an user by id",
			storageMock: StorageMock{
				getByIdMock: func() (interface{}, error) {
					return MockUsers[1], nil
				},
			},
			expectedResult: ExpectedUsers[1],
			expectedError:  nil,
		},
		{
			name: "Should get an user by id if store returns a json",
			storageMock: StorageMock{
				getByIdMock: func() (interface{}, error) {
					return mockUserJson, nil
				},
			},
			expectedResult: mockGetUser,
			expectedError:  nil,
		},
		{
			name: "Should get an error if json isnt parsable to User struct",
			storageMock: StorageMock{
				getByIdMock: func() (interface{}, error) {
					return "", nil
				},
			},
			expectedResult: structures.User{},
			expectedError: custom_errors.ServiceError{
				Code:        "InternalError",
				Description: "couldnt parse store response to go struct",
			},
		},
		{
			name: "Should get an user not found error",
			storageMock: StorageMock{
				getByIdMock: func() (interface{}, error) {
					return "", errors.New("user not found")
				},
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
				t.Fatalf("Expected error: \n %d , \n  but got error: \n %d", test.expectedError, err)
			}
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Expected \n %v, \n  but got \n %v", test.expectedResult, got)
			}
		})
	}

}

func TestCreateService(t *testing.T) {
	testTable := []struct {
		name           string
		storageMock    db.Storage
		userRequest    structures.UserRequest
		expectedResult structures.User
		expectedError  error
	}{
		{
			name: "Should create an user",
			storageMock: StorageMock{
				createMock: func() (interface{}, error) {
					return mockCreateUser, nil
				},
			},
			userRequest:    mockCreateUserRequest,
			expectedResult: mockCreateUser,
			expectedError:  nil,
		},
		{
			name: "Should return an error if id is repeated",
			storageMock: StorageMock{
				createMock: func() (interface{}, error) {
					return nil, errors.New("id already used")
				},
			},
			userRequest:    mockCreateUserRequest,
			expectedResult: structures.User{},
			expectedError: custom_errors.ServiceError{
				Code:        "IdAlreadyInUse",
				Description: "id already used",
			},
		},
		{
			name: "Should get an error if json isnt parsable to User struct",
			storageMock: StorageMock{
				createMock: func() (interface{}, error) {
					return "", nil
				},
			},
			expectedResult: structures.User{},
			expectedError: custom_errors.ServiceError{
				Code:        "InternalError",
				Description: "couldnt parse store response to go struct",
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//
			service := service.NewUserService(test.storageMock)
			got, err := service.Create(test.userRequest)
			test.expectedResult.ID = got.ID

			if !errors.Is(err, test.expectedError) {
				t.Fatalf("Expected error: \n %d , \n  but got error: \n %d", test.expectedError, err)
			}
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Expected \n %v, \n  but got \n %v", test.expectedResult, got)
			}
		})
	}

}

func TestDeleteService(t *testing.T) {
	testTable := []struct {
		name          string
		storageMock   db.Storage
		id            uuid.UUID
		expectedError error
	}{
		{
			name: "Should delete an user",
			storageMock: StorageMock{
				deleteMock: func() error {
					return nil
				},
			},
			id:            mockCreateUser.ID,
			expectedError: nil,
		},
		{
			name: "Should return error on deliting user not found",
			storageMock: StorageMock{

				deleteMock: func() error {
					return errors.New("user not found")
				},
			},
			id: uuid.UUID{},
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
			err := service.Delete(test.id)

			if !errors.Is(err, test.expectedError) {
				t.Fatalf("Expected error: \n %d , \n  but got error: \n %d", test.expectedError, err)
			}
		})
	}

}

func TestUpdateService(t *testing.T) {
	testTable := []struct {
		name           string
		storageMock    db.Storage
		expectedResult structures.User
		expectedError  error
		updateUser     structures.User
	}{
		{
			name: "Should update an user",
			storageMock: StorageMock{
				updateMock: func() (interface{}, error) {
					return mockUpdatedUser, nil
				},
			},
			expectedResult: mockUpdatedUser,
			expectedError:  nil,
			updateUser:     mockUpdatedUser,
		}, {
			name: "Should get an user not found error",
			storageMock: StorageMock{
				updateMock: func() (interface{}, error) {
					return nil, errors.New("user not found")
				},
			},
			expectedResult: structures.User{},
			expectedError: custom_errors.ServiceError{
				Code:        "NotFound",
				Description: "user not found",
			},
			updateUser: mockUpdatedUser,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//
			service := service.NewUserService(test.storageMock)
			got, err := service.Update(uuid.UUID{}, test.updateUser)
			test.expectedResult.ID = got.ID

			if !errors.Is(err, test.expectedError) {
				t.Fatalf("Expected error: \n %d , \n  but got error: \n %d", test.expectedError, err)
			}
			if !reflect.DeepEqual(got, test.expectedResult) {
				t.Errorf("Expected \n %v, \n  but got \n %v", test.expectedResult, got)
			}
		})
	}

}
