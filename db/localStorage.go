package db

import (
	"bootcam1_users/custom_errors"

	"github.com/google/uuid"
)

type localStorage struct {
	entities map[uuid.UUID]interface{}
}

func NewLocalStorage() Storage {
	local := localStorage{}

	//initialize users map
	local.entities = make(map[uuid.UUID]interface{})
	for _, user := range DefaultUsers {
		local.entities[user.ID] = user
	}

	return &local
}
func (l *localStorage) Get(id uuid.UUID) (interface{}, error) {
	entity, ok := l.entities[id]

	if !ok {
		return nil, custom_errors.Error_UserNotFound
	}

	return entity, nil
}

func (l *localStorage) GetAll() ([]interface{}, error) {
	entitiesArr := make([]interface{}, 0)
	for _, entity := range l.entities {
		entitiesArr = append(entitiesArr, entity)
	}
	return entitiesArr, nil
}

func (l *localStorage) Create(id uuid.UUID, entityToCreate interface{}) (interface{}, error) {
	//crear id aca
	_, found := l.entities[id]
	if found {
		return nil, custom_errors.Error_UuidAlreadyExists
	}

	l.entities[id] = entityToCreate

	return l.entities[id], nil

}

func (l *localStorage) Update(id uuid.UUID, entityToUpdate interface{}) (interface{}, error) {
	_, found := l.entities[id]
	if !found {
		return nil, custom_errors.Error_UserNotFound
	}

	l.entities[id] = entityToUpdate

	return l.entities[id], nil
}

func (l *localStorage) Delete(id uuid.UUID) error {
	_, found := l.entities[id]
	if !found {
		return custom_errors.Error_UserNotFound
	}

	delete(l.entities, id)

	return nil
}
