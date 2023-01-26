package storage

import (
	"github.com/altuntasfatih/car-service-backend/pkg/models"
)

type Reader interface {
	GetRepair(id string) (*models.Repair, error)
	GetAllRepairs() ([]*models.Repair, error)
}

type Writer interface {
	CreateRepair(id string, user *models.Repair) error
	UpdateRepair(id string, user *models.Repair) error
}

type Remover interface {
	DeleteRepair(id string) error
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
