package repository

import "github.com/TulioGuaraldoB/car-app/domain/entity"

type IUserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUser(userId uint) (*entity.User, error)
	GetUserByCredentials(user *entity.User) (*entity.User, error)
	CreateUser(user *entity.User) error
}
