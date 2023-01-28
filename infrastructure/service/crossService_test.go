package service_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/TulioGuaraldoB/car-app/config/env"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/TulioGuaraldoB/car-app/infrastructure/service"
	faker "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

type crossServiceTest struct {
	description            string
	expectedUserReq        *dto.UserRequest
	expectedCredentialsReq *dto.Credentials
	expectedKeycloakUrl    string
	setMocks               http.Client
}

func TestCrossServiceRegister(t *testing.T) {
	mockUserRequest := new(dto.UserRequest)
	faker.Struct(mockUserRequest)

	tests := []crossServiceTest{
		{
			description:     "Should return no error",
			expectedUserReq: mockUserRequest,
			setMocks: http.Client{
				Timeout: time.Second * 1,
			},
			expectedKeycloakUrl: "http://localhost:8082",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			t.Setenv("CROSS_ROADS_URL", testCase.expectedKeycloakUrl)
			env.GetEnvironmentVariables()
			crossService := service.NewCrossService(testCase.setMocks)

			// Act
			err := crossService.Register(testCase.expectedUserReq)
			if err != nil {
				assert.Error(t, err)
				return
			} // Assert

			assert.Nil(t, err)
		})
	}
}

func TestCrossServiceLogin(t *testing.T) {
	mockCredentialsRequest := new(dto.Credentials)
	faker.Struct(mockCredentialsRequest)

	tests := []crossServiceTest{
		{
			description:            "Should return no error",
			expectedCredentialsReq: mockCredentialsRequest,
			setMocks: http.Client{
				Timeout: time.Second * 1,
			},
			expectedKeycloakUrl: "http://localhost:8082",
		},
		{
			description: "Should return no error on admin login",
			expectedCredentialsReq: &dto.Credentials{
				Username: "admin",
				Password: "admin",
			},
			setMocks: http.Client{
				Timeout: time.Second * 1,
			},
			expectedKeycloakUrl: "http://localhost:8082",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			t.Setenv("CROSS_ROADS_URL", testCase.expectedKeycloakUrl)
			env.GetEnvironmentVariables()
			crossService := service.NewCrossService(testCase.setMocks)

			// Act
			token, err := crossService.Login(testCase.expectedCredentialsReq)
			if err != nil {
				assert.Error(t, err)
				return
			} // Assert

			assert.Nil(t, err)
			assert.NotEmpty(t, token)
		})
	}
}
