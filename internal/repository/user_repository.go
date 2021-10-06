package repository

type UserRepository interface {
	CreateUser()
	GetUsers()
	GetUser()
	DeleteUser()
}

func NewUserRepository() (UserRepository, error) {
	return nil, nil
}
