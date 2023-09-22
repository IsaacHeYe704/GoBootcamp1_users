package structures

type userManager struct {
	users map[string]User
}

func NewUserManager() (userManager, error) {
	manager := userManager{}
	//initialize users map
	manager.users = make(map[string]User)
	for _, user := range DefaultUsers {
		manager.users[user.ID] = user
	}

	return manager, nil
}

func (u *userManager) Read(uuid string) (User, error) {
	user, ok := u.users[uuid]
	if !ok {
		return User{}, Error_UserNotFound
	}
	return user, nil
}
func (u *userManager) Create(user User) (User, error) {
	_, found := u.users[user.ID]
	if found {
		return User{}, Error_UuidAlreadyExists
	}
	u.users[user.ID] = user
	return u.users[user.ID], nil
}
func (u *userManager) Update(uuid string, user User) (User, error) {
	_, found := u.users[user.ID]
	if !found {
		return User{}, Error_UserNotFound
	}
	u.users[uuid] = user
	return u.users[user.ID], nil
}
func (u *userManager) Delete(uuid string) error {
	_, found := u.users[uuid]
	if !found {
		return Error_UserNotFound
	}
	delete(u.users, uuid)
	return nil
}
