package application

import (
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/TulioGuaraldoB/car-app/infrastructure/service"
)

type ICrossApplication interface {
	Register(userRequest *dto.UserRequest) error
	Login(credentialsRequest *dto.Credentials) (*string, error)
}

type crossApplication struct {
	crossService service.ICrossService
}

func NewCrossApplication(crossService service.ICrossService) ICrossApplication {
	return &crossApplication{
		crossService: crossService,
	}
}

func (a *crossApplication) Register(userRequest *dto.UserRequest) error {
	return a.crossService.Register(userRequest)
}

func (a *crossApplication) Login(credentialsRequest *dto.Credentials) (*string, error) {
	return a.crossService.Login(credentialsRequest)
}
