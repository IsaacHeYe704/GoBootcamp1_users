package db

import (
	"github.com/google/uuid"
)

type Storage interface {
	Get(uuid.UUID) (interface{}, error)
	GetAll() ([]interface{}, error)
	Create(uuid.UUID, interface{}) (interface{}, error)
	Update(uuid.UUID, interface{}) (interface{}, error)
	Delete(uuid.UUID) error
}