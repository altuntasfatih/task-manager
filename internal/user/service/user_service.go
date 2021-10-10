package service

import (
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/storage"
	"github.com/rs/xid"
)

type UserService interface {
	CreateUser(request *models.CreateUserRequest) (*models.User, error)
	GetUser(userId string) (*models.User, error)
	DeleteUser(userId string) error
	GetUsers() ([]*models.User, error)
}

type userService struct {
	userStore storage.ReaderWriterRemover
}

func NewUserService(userStore storage.ReaderWriterRemover) (UserService, error) {
	return &userService{
		userStore: userStore,
	}, nil
}

func (u *userService) GetUsers() ([]*models.User, error) {
	return u.userStore.GetAllUsers()
}

func (u *userService) GetUser(userId string) (*models.User, error) {
	return u.userStore.GetUser(userId)
}

func (u *userService) CreateUser(request *models.CreateUserRequest) (*models.User, error) {
	guid := xid.New().String()
	newUser := models.NewUser(guid, request.Email, request.FirstName, request.LastName)
	err := u.userStore.CreateUser(guid, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, err
}

func (u *userService) DeleteUser(userId string) error {
	return u.userStore.DeleteUser(userId)
}
