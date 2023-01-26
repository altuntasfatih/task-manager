package service

import (
	"github.com/altuntasfatih/car-service-backend/pkg/models"
	"github.com/altuntasfatih/car-service-backend/pkg/storage"
	"github.com/rs/xid"
)

type RepairService interface {
	CreateRepair(request *models.CreateRepairRequest) (*models.Repair, error)
	GetRepair(userId string) (*models.Repair, error)
	DeleteRepair(userId string) error
	GetRepairs() ([]*models.Repair, error)
}

type repairService struct {
	repairStore storage.ReaderWriterRemover
}

func NewRepairService(userStore storage.ReaderWriterRemover) (RepairService, error) {
	return &repairService{
		repairStore: userStore,
	}, nil
}

func (u *repairService) GetRepairs() ([]*models.Repair, error) {
	return u.repairStore.GetAllRepairs()
}

func (u *repairService) GetRepair(repairId string) (*models.Repair, error) {
	return u.repairStore.GetRepair(repairId)
}

func (u *repairService) CreateRepair(request *models.CreateRepairRequest) (*models.Repair, error) {
	guid := xid.New().String()
	newUser := models.NewRepair(guid, request.Email, request.FirstName, request.LastName)
	err := u.repairStore.CreateRepair(guid, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, err
}

func (u *repairService) DeleteRepair(repairId string) error {
	return u.repairStore.DeleteRepair(repairId)
}
