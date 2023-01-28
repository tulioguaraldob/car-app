package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TulioGuaraldoB/car-app/config/env"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
)

type ICrossService interface {
	Register(userRequest *dto.UserRequest) error
	Login(credentialsRequest *dto.Credentials) (*string, error)
}

type crossService struct {
	httpClient http.Client
}

func NewCrossService(httpClient http.Client) ICrossService {
	return &crossService{
		httpClient: httpClient,
	}
}

func (s *crossService) Register(userRequest *dto.UserRequest) error {
	crossUrl := fmt.Sprintf("%s/api/v1/user/register", env.Env.CrossroadsUrl)
	userRequestBuffer, err := BindUserRequest(userRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, crossUrl, userRequestBuffer)
	if err != nil {
		return err
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		errMessage := fmt.Errorf("failed to register user on cross-roads. %v", res)
		return errMessage
	}

	return nil
}

func (s *crossService) Login(credentialsRequest *dto.Credentials) (*string, error) {
	crossUrl := fmt.Sprintf("%s/api/v1/user/login", env.Env.CrossroadsUrl)
	userRequestBuffer, err := BindUserRequest(credentialsRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, crossUrl, userRequestBuffer)
	if err != nil {
		return nil, err
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		errMessage := fmt.Errorf("failed to login. %v", res)
		return nil, errMessage
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	parsedResBody := fmt.Sprintf("%v", string(resBody))
	return &parsedResBody, nil
}

func BindUserRequest(userRequest any) (io.Reader, error) {
	userRequestJSON, err := json.Marshal(userRequest)
	if err != nil {
		return nil, err
	}

	userRequestBuffer := bytes.NewBuffer(userRequestJSON)
	return userRequestBuffer, nil
}
