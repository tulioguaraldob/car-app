package persistence_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence"
	faker "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

type carRepositoryTest struct {
	description   string
	expectedQuery string
}

func TestGetAllCarsRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `cars` WHERE `cars`.`deleted_at` IS NULL",
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `cars` WHERE `cars``deleted_at`IS NULL",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			mockCar := entity.Car{}
			faker.Struct(&mockCar)

			rows := sqlmock.NewRows([]string{
				"ID",
				"Name",
				"Brand",
				"LicensePlate",
				"Year",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockCar.ID,
				mockCar.Name,
				mockCar.Brand,
				mockCar.LicensePlate,
				mockCar.Year,
				mockCar.CreatedAt,
				mockCar.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			dbMock, sqlMock := mockDb()

			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			_, err := carRepository.GetAllCars()
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
