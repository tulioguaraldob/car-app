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

	// Handlers
	userHandler := interfaces.NewUserHandler(userApplication)

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
		}
	}

	return router
}
