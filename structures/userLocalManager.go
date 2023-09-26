package structures

import (
	"bootcam1_users/custom_errors"

	"github.com/google/uuid"
)

type userLocalManager struct {
	users map[uuid.UUID]User
}

func NewUserManager() userLocalManager {
	manager := userLocalManager{}

	//initialize users map
	manager.users = make(map[uuid.UUID]User)
	for _, user := range DefaultUsers {
		manager.users[user.ID] = user
	}

	return manager
}

func (u *userLocalManager) Get(uuid uuid.UUID) (User, error) {
	user, ok := u.users[uuid]
	if !ok {
		return User{}, custom_errors.Error_UserNotFound
	}

	return user, nil
}
func (u *userLocalManager) GetAll() (map[uuid.UUID]User, error) {

	return u.users, nil
}
func (u *userLocalManager) Create(user User) (User, error) {
	//crear id aca
	_, found := u.users[user.ID]
	if found {
		return User{}, custom_errors.Error_UuidAlreadyExists
	}

	u.users[user.ID] = user

	return u.users[user.ID], nil
}
func (u *userLocalManager) Update(uuid uuid.UUID, user User) (User, error) {
	_, found := u.users[user.ID]
	if !found {
		return User{}, custom_errors.Error_UserNotFound
	}

	u.users[uuid] = user

	return u.users[user.ID], nil
}
func (u *userLocalManager) Delete(uuid uuid.UUID) error {
	_, found := u.users[uuid]
	if !found {
		return custom_errors.Error_UserNotFound
	}

	delete(u.users, uuid)

	return nil
}
