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

type carApplicationTest struct {
	description string
	setMocks    func(*mock.MockICarRepository)
}

func TestGetAllCarsApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	mockCars := []entity.Car{}
	mockCars = append(mockCars, mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					GetAllCars().
					Return(mockCars, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			carApplication := application.NewCarApplication(mcr)
			_, err := carApplication.GetAllCars()
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetCarByIdApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					GetCarById(mockCar.ID).
					Return(&mockCar, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			userApplication := application.NewCarApplication(mcr)
			_, err := userApplication.GetCarById(mockCar.ID)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetCarsByBrandApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	mockCars := []entity.Car{}
	mockCars = append(mockCars, mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					GetCarsByBrand(mockCar.Brand).
					Return(mockCars, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			userApplication := application.NewCarApplication(mcr)
			_, err := userApplication.GetCarsByBrand(mockCar.Brand)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetCarByLicensePlateApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					GetCarByLicensePlate(mockCar.LicensePlate).
					Return(&mockCar, nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			userApplication := application.NewCarApplication(mcr)
			_, err := userApplication.GetCarByLicensePlate(mockCar.LicensePlate)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestCreateCarApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					CreateCar(&mockCar).
					Return(nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			userApplication := application.NewCarApplication(mcr)
			err := userApplication.CreateCar(&mockCar)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestDeleteCarApplication(t *testing.T) {
	mockCar := entity.Car{}
	faker.Struct(&mockCar)

	tests := []carApplicationTest{
		{
			description: "Should return no error on get user",
			setMocks: func(mir *mock.MockICarRepository) {
				mir.EXPECT().
					DeleteCar(mockCar.ID).
					Return(nil)
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			controller := gomock.NewController(t)
			defer controller.Finish()

			mcr := mock.NewMockICarRepository(controller)
			testCase.setMocks(mcr)

			// Act
			userApplication := application.NewCarApplication(mcr)
			err := userApplication.DeleteCar(mockCar.ID)

			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
