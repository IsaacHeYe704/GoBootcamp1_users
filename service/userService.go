package service

import (
	"bootcam1_users/structures"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	storage structures.Storage
}

func NewUserService(storage structures.Storage) UserService {
	return UserService{storage: storage}
}

func (us *UserService) Get(uuid uuid.UUID) (structures.User, error) {

	slog.Info("Getting user with ", "id", uuid.String())
	return us.storage.Get(uuid)
}
func (us *UserService) GetAll() ([]structures.User, error) {
	slog.Info("Getting all users ")
	return us.storage.GetAll()
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
	user, err := us.storage.Create(userParsed)
	return user, err
}
func (us *UserService) Update(uuid uuid.UUID, user structures.User) (structures.User, error) {
	slog.Info("Updating user with ", "id", uuid.String(), ", new data: ", fmt.Sprint(user))

	return us.storage.Update(uuid, user)
}
func (us *UserService) Delete(uuid uuid.UUID) error {
	slog.Info("Deleting user with ", "id", uuid.String())
	return us.storage.Delete(uuid)
}
