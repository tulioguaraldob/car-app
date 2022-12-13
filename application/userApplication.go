package application

import (
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/domain/repository"
)

type IUserApplication interface {
	GetAllUsers() ([]entity.User, error)
	GetUser(userId uint) (*entity.User, error)
	GetUserByCredentials(user *entity.User) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userApplication struct {
	userRepository repository.IUserRepository
}

func NewUserApplication(userRepository repository.IUserRepository) IUserApplication {
	return &userApplication{
		userRepository: userRepository,
	}
}

var _ IUserApplication = &userApplication{}

func (a *userApplication) GetAllUsers() ([]entity.User, error) {
	return a.userRepository.GetAllUsers()
}

func (a *userApplication) GetUser(userId uint) (*entity.User, error) {
	return a.userRepository.GetUser(userId)
}

func (a *userApplication) GetUserByCredentials(user *entity.User) (*entity.User, error) {
	return a.userRepository.GetUserByCredentials(user)
}

func (a *userApplication) CreateUser(user *entity.User) error {
	return a.userRepository.CreateUser(user)
}
