package persistence

import (
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	users := []entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUser(userId uint) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.First(&user, &userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByCredentials() (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}
