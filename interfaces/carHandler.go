package interfaces

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type carHandler struct {
	carApplication application.ICarApplication
}

func NewCarHandler(carApplication application.ICarApplication) *carHandler {
	return &carHandler{
		carApplication: carApplication,
	}
}

func (h *carHandler) GetAllCars(ctx *gin.Context) {
	cars, err := h.carApplication.GetAllCars()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if len(cars) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})

		return
	}

	carsResponse := []dto.CarResponse{}
	for _, car := range cars {
		carResponse := dto.CarToResponse(&car)
		carsResponse = append(carsResponse, *carResponse)
	}

	ctx.JSON(http.StatusOK, carsResponse)
}

func (h *carHandler) GetCarById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	car, err := h.carApplication.GetCarById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	carResponse := dto.CarToResponse(car)
	ctx.JSON(http.StatusOK, carResponse)
}

func (h *carHandler) GetCarsByBrand(ctx *gin.Context) {
	brand := ctx.Param("brand")
	cars, err := h.carApplication.GetCarsByBrand(brand)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	if len(cars) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})

		return
	}

	carsResponse := []dto.CarResponse{}
	for _, car := range cars {
		carResponse := dto.CarToResponse(&car)
		carsResponse = append(carsResponse, *carResponse)
	}

	ctx.JSON(http.StatusOK, carsResponse)
}

func (h *carHandler) GetCarByLicensePlate(ctx *gin.Context) {
	licensePlate := ctx.Param("licensePlate")
	car, err := h.carApplication.GetCarByLicensePlate(licensePlate)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	carResponse := dto.CarToResponse(car)
	ctx.JSON(http.StatusOK, carResponse)
}

func (h *carHandler) CreateCar(ctx *gin.Context) {
	carRequest := new(dto.CarRequest)
	if err := ctx.ShouldBindJSON(carRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	car := dto.RequestToCar(carRequest)
	if err := h.carApplication.CreateCar(car); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "car inserted successfully!",
		"car":     carRequest,
	})
}
