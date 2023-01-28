package routes

import (
	"net/http"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence/db"
	"github.com/TulioGuaraldoB/car-app/infrastructure/service"
	"github.com/TulioGuaraldoB/car-app/interfaces"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	// Services
	httpClient := new(http.Client)
	repositories := db.NewRepositories()

	crossService := service.NewCrossService(*httpClient)

	// Application
	userApplication := application.NewUserApplication(repositories.IUserRepository)
	carApplication := application.NewCarApplication(repositories.ICarRepository)
	crossApplication := application.NewCrossApplication(crossService)

	// Handlers
	userHandler := interfaces.NewUserHandler(userApplication)
	carHandler := interfaces.NewCarHandler(carApplication)
	crossHandler := interfaces.NewCrossHandler(crossApplication)

	router := gin.Default()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			user := v1.Group("user")
			{
				user.GET("", userHandler.GetAllUsers)
				user.GET(":id", userHandler.GetUserById)
				user.POST("login", userHandler.Login)
				user.POST("register", userHandler.Register)
			}

			car := v1.Group("car")
			{
				car.GET("", carHandler.GetAllCars)
				car.GET(":id", carHandler.GetCarById)
				car.GET("brand/:brand", carHandler.GetCarsByBrand)
				car.GET("license/:licensePlate", carHandler.GetCarByLicensePlate)
				car.POST("", carHandler.CreateCar)
			}

			cross := v1.Group("cross")
			{
				cross.POST("register", crossHandler.Register)
				cross.POST("login", crossHandler.Login)
			}
		}
	}

	return router
}
