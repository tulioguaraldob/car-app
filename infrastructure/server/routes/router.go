package routes

import (
	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence/db"
	"github.com/TulioGuaraldoB/car-app/interfaces"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	// Services
	repositories := db.NewRepositories()

	// Application
	userApplication := application.NewUserApplication(repositories.IUserRepository)
	carApplication := application.NewCarApplication(repositories.ICarRepository)

	// Handlers
	userHandler := interfaces.NewUserHandler(userApplication)
	carHandler := interfaces.NewCarHandler(carApplication)

	router := gin.Default()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			user := v1.Group("user")
			{
				user.GET("", userHandler.GetAllUsers)
				user.GET(":id", userHandler.GetUserById)
			}

			car := v1.Group("car")
			{
				car.GET("", carHandler.GetAllCars)
				car.GET(":id", carHandler.GetCarById)
				car.GET("brand/:brand", carHandler.GetCarsByBrand)
				car.GET("license/:licensePlate", carHandler.GetCarByLicensePlate)
				car.POST("", carHandler.CreateCar)
			}
		}
	}

	return router
}
