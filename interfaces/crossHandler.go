package interfaces

import (
	"net/http"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/gin-gonic/gin"
)

type crossHandler struct {
	crossApplication application.ICrossApplication
}

func NewCrossHandler(crossApplication application.ICrossApplication) *crossHandler {
	return &crossHandler{
		crossApplication: crossApplication,
	}
}

func (h *crossHandler) Register(ctx *gin.Context) {
	userRequest := new(dto.UserRequest)
	if err := ctx.ShouldBindJSON(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := h.crossApplication.Register(userRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully!",
		"user":    userRequest,
	})
}

func (h *crossHandler) Login(ctx *gin.Context) {
	credentialsRequest := new(dto.Credentials)
	if err := ctx.ShouldBindJSON(credentialsRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	token, err := h.crossApplication.Login(credentialsRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, token)
}
