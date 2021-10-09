package store

import (
	"github.com/altuntasfatih/task-manager/pkg/models"
)


type Reader interface {
	GetUser(id string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}

type Writer interface {
	CreateUser(id string, user *models.User) error
	UpdateUser(id string, user *models.User) error
}

type Remover interface {
	DeleteUser(id string) error
}

type ReaderWriter interface {
	Reader
	Writer
}

type ReaderWriterRemover interface {
	Reader
	Writer
	Remover
}
