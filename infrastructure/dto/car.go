package dto

import (
	"strconv"
	"time"

	"github.com/TulioGuaraldoB/car-app/domain/entity"
)

type CarResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	LicensePlate string `json:"license_plate"`
	Year         int64  `json:"year"`
}

type CarRequest struct {
	Name         string `json:"name"`
	Brand        string `json:"brand"`
	LicensePlate string `json:"license_plate"`
	Year         int64  `json:"year"`
}

func CarToResponse(car *entity.Car) *CarResponse {
	parsedStringYear := car.Year.Format("2006")
	parsedYear, _ := strconv.Atoi(parsedStringYear)

	return &CarResponse{
		ID:           car.ID,
		Name:         car.Name,
		Brand:        car.Brand,
		LicensePlate: car.LicensePlate,
		Year:         int64(parsedYear),
	}
}

func RequestToCar(carReq *CarRequest) *entity.Car {
	return &entity.Car{
		Name:         carReq.Name,
		Brand:        carReq.Brand,
		LicensePlate: carReq.LicensePlate,
		Year:         time.Date(int(carReq.Year), time.January, 1, 0, 0, 0, 0, time.Local),
	}
}
