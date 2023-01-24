package persistence_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence/mock"
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
			dbMock, sqlMock := mock.MockDatabase()

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

func TestGetCarByIdRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `cars` WHERE `cars`.`id` = ? AND `cars`.`deleted_at` IS NULL ORDER BY `cars`.`id` LIMIT 1",
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `cars` WHERE `cars``id` = ?AND `cars`.`deleted_at` IS NULL ORDER BY `cars`.`id` LIMIT 1",
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
			dbMock, sqlMock := mock.MockDatabase()

			sqlMock.ExpectQuery(query).WithArgs(mockCar.ID).WillReturnRows(rows)

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			_, err := carRepository.GetCarById(mockCar.ID)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetCarsByBrandRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `cars` WHERE brand = ? AND `cars`.`deleted_at` IS NULL",
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `cars` WHERE brand = ?AND `cars``deleted_at` IS NULL",
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
			dbMock, sqlMock := mock.MockDatabase()

			sqlMock.ExpectQuery(query).WithArgs(mockCar.Brand).WillReturnRows(rows)

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			_, err := carRepository.GetCarsByBrand(mockCar.Brand)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestGetCarByLicensePlateRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "SELECT * FROM `cars` WHERE license_plate = ? AND `cars`.`deleted_at` IS NULL ORDER BY `cars`.`id` LIMIT 1",
		},
		{
			description:   "Should return error",
			expectedQuery: "SELECT * FROM `cars` WHERE license_plate = ?AND `cars``deleted_at` IS NULL ORDER BY `cars`.`id` LIMIT 1",
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
			dbMock, sqlMock := mock.MockDatabase()

			sqlMock.ExpectQuery(query).WithArgs(mockCar.LicensePlate).WillReturnRows(rows)

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			_, err := carRepository.GetCarByLicensePlate(mockCar.LicensePlate)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestCreateCarRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "INSERT INTO `cars` (`name`,`brand`,`license_plate`,`year`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)",
		},
		{
			description:   "Should return error",
			expectedQuery: "INSERT INTO `cars` (`name`,`brand`,`license_plate``year`,`created_at`,`updated_at`,`deleted_at`) VALUES(?,?,?,?,?,?,?)",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			mockCar := entity.Car{}
			faker.Struct(&mockCar)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			dbMock, sqlMock := mock.MockDatabase()

			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(query).
				WillReturnResult(sqlmock.NewResult(1, 1))
			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			err := carRepository.CreateCar(&mockCar)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestDeleteCarRepository(t *testing.T) {
	tests := []carRepositoryTest{
		{
			description:   "Should return no error",
			expectedQuery: "UPDATE `cars` SET `deleted_at`=? WHERE id = ? AND `cars`.`deleted_at` IS NULL",
		},
		{
			description:   "Should return error",
			expectedQuery: "UPDATE `cars` SET `deleted_at`=? WHERE id = ?AND `cars``deleted_at` IS NULL",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			mockCar := entity.Car{}
			faker.Struct(&mockCar)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			dbMock, sqlMock := mock.MockDatabase()

			sqlMock.ExpectBegin()
			sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			sqlMock.ExpectCommit()
			sqlMock.ExpectClose()

			// Act
			carRepository := persistence.NewCarRepository(dbMock)
			err := carRepository.DeleteCar(mockCar.ID)
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
