package service

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/db"
	"bootcam1_users/structures"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	storage db.Storage
}

func NewUserService(storage db.Storage) UserService {
	return UserService{storage: storage}
}

func (us *UserService) Get(uuid uuid.UUID) (structures.User, error) {

	slog.Info("Getting user with ", "id", uuid.String())
	response, err := us.storage.Get(uuid)
	if err != nil {
		return structures.User{}, custom_errors.ServiceError{
			Code:        "NotFound",
			Description: err.Error(),
		}
	}

	//asert user
	user, ok := response.(structures.User)
	if !ok {
		//user is not a user struct
		// try to parse it
		user, err = parseUser(fmt.Sprint(response))
		if err != nil {
			return structures.User{}, custom_errors.ServiceError{
				Code:        "InternalError",
				Description: "couldnt parse store response to go struct",
			}
		}
	}

	return user, nil

}
func (us *UserService) GetAll() ([]structures.User, error) {
	slog.Info("Getting all users ")

	response, err := us.storage.GetAll()
	if err != nil {
		return nil, custom_errors.ServiceError{
			Code:        "ConectionError",
			Description: "connection refused",
		}
	}
	users := make([]structures.User, 0)
	//since storage returns []interface{} we should assert or parse that into user Struct so we can return []structures.User
	for _, v := range response {
		//check if the item can be asserted to User struct
		if val, ok := v.(structures.User); ok {
			users = append(users, val)
		} else {
			//try to parse the item into the user struct
			user, err := parseUser(fmt.Sprint(v))
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
	}

	return users, err
}
func (us *UserService) Create(userRequest structures.UserRequest) (structures.User, error) {
	slog.Info("Creating ", "user", fmt.Sprint(userRequest))
	//create uuid
	newUuid := uuid.New()
	userParsed := structures.User{
		ID:       newUuid,
		Name:     userRequest.Name,
		LastName: userRequest.LastName,
		Email:    userRequest.LastName,
		Active:   userRequest.Active,
		Address:  userRequest.Address,
	}
	response, err := us.storage.Create(newUuid, userParsed)
	if err != nil {
		return structures.User{}, custom_errors.ServiceError{
			Code:        "IdAlreadyInUse",
			Description: err.Error(),
		}
	}
	//asert user
	user, ok := response.(structures.User)
	if !ok {
		//user is not a user struct
		// try to parse it
		user, err = parseUser(fmt.Sprint(response))
		if err != nil {
			return structures.User{}, err
		}
	}
	return user, err
}
func (us *UserService) Update(uuid uuid.UUID, user structures.User) (structures.User, error) {
	slog.Info("Updating user with ", "id", uuid.String(), ", new data: ", fmt.Sprint(user))
	response, err := us.storage.Update(uuid, user)
	userUpdated, _ := response.(structures.User)
	return userUpdated, err
}
func (us *UserService) Delete(uuid uuid.UUID) error {
	slog.Info("Deleting user with ", "id", uuid.String())

	err := us.storage.Delete(uuid)
	if err != nil {
		return custom_errors.ServiceError{
			Code:        "NotFound",
			Description: err.Error(),
		}
	}
	return nil
}

func parseUser(jsonVal string) (structures.User, error) {
	data := structures.User{}
	err := json.Unmarshal([]byte(jsonVal), &data)
	if err != nil {
		return structures.User{}, err
	}
	return data, nil
}
