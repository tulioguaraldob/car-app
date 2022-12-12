package persistence

import (
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/domain/repository"
	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) repository.ICarRepository {
	return &carRepository{
		db: db,
	}
}

func (r *carRepository) GetAllCars() ([]entity.Car, error) {
	cars := []entity.Car{}
	if err := r.db.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *carRepository) GetCarById(carId uint) (*entity.Car, error) {
	car := entity.Car{}
	if err := r.db.First(&car, &carId).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *carRepository) GetCarsByBrand(brand string) ([]entity.Car, error) {
	cars := []entity.Car{}
	if err := r.db.Where("brand = ?", &brand).Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *carRepository) GetCarByLicensePlate(licensePlate string) (*entity.Car, error) {
	car := entity.Car{}
	if err := r.db.Where("license_plate = ?", &licensePlate).First(&car).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *carRepository) CreateCar(car *entity.Car) error {
	return r.db.Create(car).Error
}

func (r *carRepository) DeleteCar(carId uint) error {
	return r.db.Where("id = ?", &carId).Delete(&entity.Car{}).Error
}
