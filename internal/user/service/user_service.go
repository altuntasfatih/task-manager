package service

import (
	"github.com/altuntasfatih/task-manager/internal/repository"
	"github.com/altuntasfatih/task-manager/internal/user/models"
)

type UserService interface {
	GetUsers() ([]*models.User, error)
	GetUser() (*models.User, error)
	CreateUser() (*models.User, error)
	DeleteUser() error
}
type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) (UserService, error) {
	return &userService{
		userRepository: userRepository,
	}, nil
}

func (u *userService) GetUsers() ([]*models.User, error) {
	return nil, nil
}
func (u *userService) GetUser() (*models.User, error) {
	return nil, nil
}
func (u *userService) CreateUser() (*models.User, error) {
	return nil, nil
}
func (u *userService) DeleteUser() error {
	return nil
}
