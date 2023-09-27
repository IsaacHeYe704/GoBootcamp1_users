package structures

import (
	"github.com/google/uuid"
)

type Storage interface {
	Get(uuid.UUID) (User, error)
	GetAll() ([]User, error)
	Create(User) (User, error)
	Update(uuid.UUID, User) (User, error)
	Delete(uuid.UUID) error
}
