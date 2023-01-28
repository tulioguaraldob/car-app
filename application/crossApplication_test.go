package application_test

import (
	"testing"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/infrastructure/dto"
	"github.com/TulioGuaraldoB/car-app/infrastructure/service/mock"
	faker "github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type crossApplicationTest struct {
	description string
	setMocks    func(*mock.MockICrossService)
}

func TestRegisterCrossApplication(t *testing.T) {
	mockUserRequest := dto.UserRequest{}
	faker.Struct(&mockUserRequest)

	tests := []crossApplicationTest{
		{
			description: "Should return no error",
			setMocks: func(mic *mock.MockICrossService) {
				mic.EXPECT().
					Register(&mockUserRequest).
					Return(nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcs := mock.NewMockICrossService(controller)
			testCase.setMocks(mcs)

			// Act
			crossApplication := application.NewCrossApplication(mcs)
			err := crossApplication.Register(&mockUserRequest)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestLoginCrossApplication(t *testing.T) {
	mockCredentialsRequest := dto.Credentials{}
	faker.Struct(&mockCredentialsRequest)

	mockToken := new(string)
	*mockToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJCUDFOVHhRb1daVlU0MkFQR2ZNMG9xV0NBRVhCWTFSZDc1TVZ0ZFBzZXlvIn0.eyJleHAiOjE2NzQ4Nzk0NDYsImlhdCI6MTY3NDg3OTM4NiwianRpIjoiMjk0NTUyYTctYWZhMy00YWJmLThlOTgtM2FhMjFlNWJmMjhkIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDAwL3JlYWxtcy9tYXN0ZXIiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiOGM4MGY5YWQtYTA4Ni00Yjg3LWJiZDYtNDJiMjhmZWUxOWQxIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiY3Jvc3Mtcm9hZHMiLCJzZXNzaW9uX3N0YXRlIjoiODMxZmE4NDAtMjkxMy00MjE1LTlmMDItMTYxYTQ3NTczNGUxIiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJkZWZhdWx0LXJvbGVzLW1hc3RlciIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiI4MzFmYTg0MC0yOTEzLTQyMTUtOWYwMi0xNjFhNDc1NzM0ZTEiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJCaWxibyBCYWdnaW5ncyIsInByZWZlcnJlZF91c2VybmFtZSI6ImJpbGJvLmIiLCJnaXZlbl9uYW1lIjoiQmlsYm8iLCJmYW1pbHlfbmFtZSI6IkJhZ2dpbmdzIiwiZW1haWwiOiJiaWxiby5iYWdnaW5zQGdtYWlsLmNvbSJ9.CO4NuH7NlvZQFdnUYL5ExmVUxw4l8Z_dtRsDbKC7C-8WHvxmU1sXdmakYLp9_J3KTq-nR8limG5x28jYHVilQ___hmuB-mI5nNGIvxUKuI2EgrhEtVQ64WRZGV5Lq-jN7Sv0B_cnweMA-3OLhH5HfohQ-ysuvLwm7hHwfhPnPVBjvQ0awbByXIPpAOdwgXruQuNWO4cwv_c4_qxjN84LuoNI4CoKS2pLOVXvDY8SZ75gxODMYsIgIgweHoDXy0fc88k7xMYszzIHjLHp4LEA0kXFRMXUXRqhf9_08S2GcK8m8H21LPvjtjdiXXXdKxk34jvYJ4-gQAd_emvzS0zA3g"

	tests := []crossApplicationTest{
		{
			description: "Should return no error",
			setMocks: func(mic *mock.MockICrossService) {
				mic.EXPECT().
					Login(&mockCredentialsRequest).
					Return(mockToken, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcs := mock.NewMockICrossService(controller)
			testCase.setMocks(mcs)

			// Act
			crossApplication := application.NewCrossApplication(mcs)
			token, err := crossApplication.Login(&mockCredentialsRequest)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotEmpty(t, token)
		})
	}
}
