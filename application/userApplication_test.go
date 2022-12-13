package application_test

import (
	"testing"

	"github.com/TulioGuaraldoB/car-app/application"
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/domain/repository/mock"
	faker "github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type userApplicationTest struct {
	description    string
	setMocks       func(*mock.MockIUserRepository)
	expectedUserId uint
}

func TestGetAllUsersApplication(t *testing.T) {
	mockUser := entity.User{}
	faker.Struct(&mockUser)

	mockUsers := []entity.User{}
	mockUsers = append(mockUsers, mockUser)

	tests := []userApplicationTest{
		{
			description: "Should return no error on get all users",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetAllUsers().
					Return(mockUsers, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(mur)

			// Act
			userApplication := application.NewUserApplication(mur)
			_, err := userApplication.GetAllUsers()

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetUserApplication(t *testing.T) {
	mockUser := entity.User{}
	faker.Struct(&mockUser)

	tests := []userApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUser(mockUser.ID).
					Return(&mockUser, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(mur)

			// Act
			userApplication := application.NewUserApplication(mur)
			_, err := userApplication.GetUser(mockUser.ID)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetUserByCredentialsApplication(t *testing.T) {
	mockUser := entity.User{}
	faker.Struct(&mockUser)

	tests := []userApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					GetUserByCredentials(&mockUser).
					Return(&mockUser, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(mur)

			// Act
			userApplication := application.NewUserApplication(mur)
			_, err := userApplication.GetUserByCredentials(&mockUser)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestCreateUserApplication(t *testing.T) {
	mockUser := entity.User{}
	faker.Struct(&mockUser)

	tests := []userApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockIUserRepository) {
				mir.EXPECT().
					CreateUser(&mockUser).
					Return(nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mur := mock.NewMockIUserRepository(controller)
			testCase.setMocks(mur)

			// Act
			userApplication := application.NewUserApplication(mur)
			err := userApplication.CreateUser(&mockUser)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
