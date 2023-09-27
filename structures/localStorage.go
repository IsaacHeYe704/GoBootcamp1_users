package structures

import (
	"bootcam1_users/custom_errors"

	"github.com/google/uuid"
)

type localStorage struct {
	users map[uuid.UUID]User
}

func NewLocalStorage() Storage {
	local := localStorage{}

	//initialize users map
	local.users = make(map[uuid.UUID]User)
	for _, user := range DefaultUsers {
		local.users[user.ID] = user
	}

	return &local
}

func (l *localStorage) Get(uuid uuid.UUID) (User, error) {
	user, ok := l.users[uuid]

	if !ok {
		return User{}, custom_errors.Error_UserNotFound
	}

	return user, nil
}
func (l *localStorage) GetAll() ([]User, error) {
	users := make([]User, 0)
	for _, val := range l.users {
		users = append(users, val)
	}
	return users, nil
}
func (l *localStorage) Create(user User) (User, error) {
	//crear id aca
	_, found := l.users[user.ID]
	if found {
		return User{}, custom_errors.Error_UuidAlreadyExists
	}

	l.users[user.ID] = user

	return l.users[user.ID], nil
}
func (l *localStorage) Update(uuid uuid.UUID, user User) (User, error) {
	_, found := l.users[user.ID]
	if !found {
		return User{}, custom_errors.Error_UserNotFound
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
