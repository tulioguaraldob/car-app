package application

import (
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/domain/repository"
)

type ICarApplication interface {
	GetAllCars() ([]entity.Car, error)
	GetCarById(carId uint) (*entity.Car, error)
	GetCarsByBrand(brand string) ([]entity.Car, error)
	GetCarByLicensePlate(licensePlate string) (*entity.Car, error)
	CreateCar(car *entity.Car) error
	DeleteCar(carId uint) error
}

type carApplication struct {
	carRepository repository.ICarRepository
}

func NewCarApplication(carRepository repository.ICarRepository) ICarApplication {
	return &carApplication{
		carRepository: carRepository,
	}
}

func (a *carApplication) GetAllCars() ([]entity.Car, error) {
	return a.carRepository.GetAllCars()
}

func (a *carApplication) GetCarById(carId uint) (*entity.Car, error) {
	return a.carRepository.GetCarById(carId)
}

func (a *carApplication) GetCarsByBrand(brand string) ([]entity.Car, error) {
	return a.carRepository.GetCarsByBrand(brand)
}

func (a *carApplication) GetCarByLicensePlate(licensePlate string) (*entity.Car, error) {
	return a.carRepository.GetCarByLicensePlate(licensePlate)
}

func (a *carApplication) CreateCar(car *entity.Car) error {
	return a.carRepository.CreateCar(car)
}

func (a *carApplication) DeleteCar(carId uint) error {
	return a.carRepository.DeleteCar(carId)
}
