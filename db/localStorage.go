package db

import (
	"bootcam1_users/custom_errors"
	"bootcam1_users/structures"

	"github.com/google/uuid"
)

type localStorage struct {
	users map[uuid.UUID]structures.User
}

func NewLocalStorage() Storage {
	local := localStorage{}

	//initialize users map
	local.users = make(map[uuid.UUID]structures.User)
	for _, user := range DefaultUsers {
		local.users[user.ID] = user
	}

	return &local
}

func (l *localStorage) Get(uuid uuid.UUID) (structures.User, error) {
	user, ok := l.users[uuid]

	if !ok {
		return structures.User{}, custom_errors.Error_UserNotFound
	}

	return user, nil
}
func (l *localStorage) GetAll() ([]structures.User, error) {
	users := make([]structures.User, 0)
	for _, val := range l.users {
		users = append(users, val)
	}
	return users, nil
}
func (l *localStorage) Create(user structures.User) (structures.User, error) {
	//crear id aca
	_, found := l.users[user.ID]
	if found {
		return structures.User{}, custom_errors.Error_UuidAlreadyExists
	}

	l.users[user.ID] = user

	return l.users[user.ID], nil
}
func (l *localStorage) Update(uuid uuid.UUID, user structures.User) (structures.User, error) {
	_, found := l.users[user.ID]
	if !found {
		return structures.User{}, custom_errors.Error_UserNotFound
	}

	l.users[uuid] = user

	return l.users[user.ID], nil
}
func (l *localStorage) Delete(uuid uuid.UUID) error {
	_, found := l.users[uuid]
	if !found {
		return custom_errors.Error_UserNotFound
	}

	delete(l.users, uuid)

	return nil
}
