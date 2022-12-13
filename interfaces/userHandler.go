package interfaces

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	userApplication application.IUserApplication
}

func NewUserHandler(userApplication application.IUserApplication) *userHandler {
	return &userHandler{
		userApplication: userApplication,
	}
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userApplication.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if len(users) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})

		return
	}

	usersRes := []dto.UserResponse{}
	for _, user := range users {
		userRes := dto.UserToResponse(&user)
		usersRes = append(usersRes, *userRes)
	}

	ctx.JSON(http.StatusOK, usersRes)
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := h.userApplication.GetUser(uint(id))
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
	}

	userRes := dto.UserToResponse(user)
	ctx.JSON(http.StatusOK, userRes)
}

func (h *userHandler) Login(ctx *gin.Context) {
	credentialsReq := new(dto.Credentials)
	if err := ctx.ShouldBindJSON(credentialsReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userReq := dto.CredentialsToUser(credentialsReq)
	user, err := h.userApplication.GetUserByCredentials(userReq)
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

	ctx.JSON(http.StatusOK, user)
}
