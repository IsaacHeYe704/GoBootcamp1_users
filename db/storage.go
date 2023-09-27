package db

import (
	"bootcam1_users/structures"

	"github.com/google/uuid"
)

type Storage interface {
	Get(uuid.UUID) (structures.User, error)
	GetAll() ([]structures.User, error)
	Create(structures.User) (structures.User, error)
	Update(uuid.UUID, structures.User) (structures.User, error)
	Delete(uuid.UUID) error
}
