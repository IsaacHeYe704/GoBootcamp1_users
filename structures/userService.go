package structures

import (
	"log/slog"

	"github.com/google/uuid"
)

type userService struct {
	storage Storage
}

func NewUserService(storage Storage) userService {
	return userService{storage: storage}
}

func (us *userService) Get(uuid uuid.UUID) (User, error) {

	slog.Info("Getting user with id ", uuid)
	return us.storage.Get(uuid)
}
func (us *userService) GetAll() (map[uuid.UUID]User, error) {
	slog.Info("Getting all users ")
	return us.storage.GetAll()
}
func (us *userService) Create(user User) (User, error) {
	slog.Info("Creating user", user)

	return us.storage.Create(user)
}
func (us *userService) Update(uuid uuid.UUID, user User) (User, error) {
	slog.Info("Updating user with id: ", uuid, ", new data: ", user)

	return us.storage.Update(uuid, user)
}
func (us *userService) Delete(uuid uuid.UUID) error {
	slog.Info("Deleting user with id: ", uuid)
	return us.storage.Delete(uuid)
}
