package repository

import "github.com/TulioGuaraldoB/car-app/domain/entity"

type ICarRepository interface {
	GetAllCars() ([]entity.Car, error)
	GetCarById(carId uint) (*entity.Car, error)
	GetCarsByBrand(brand string) ([]entity.Car, error)
	GetCarByLicensePlate(licensePlate string) (*entity.Car, error)
	CreateCar(car *entity.Car) error
	DeleteCar(carId uint) error
}
