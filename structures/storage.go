package structures

import (
	"github.com/google/uuid"
)

type Storage interface {
	NewUserManager() Storage
	Get(uuid.UUID) (User, error)
	GetAll() (map[uuid.UUID]User, error)
	Create(User) (User, error)
	Update(uuid.UUID, User) (User, error)
	Delete(uuid.UUID) error
}
