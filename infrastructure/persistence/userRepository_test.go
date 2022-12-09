package persistence_test

import (
	"log"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/TulioGuaraldoB/car-app/domain/entity"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence"
	faker "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type userRepositoryTest struct {
	description   string
	expectedQuery string
}

func mockDb() (*gorm.DB, sqlmock.Sqlmock) {
	mock, sqlMock, _ := sqlmock.New()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      mock,
		SkipInitializeWithVersion: true,
	})

	dbMock, err := gorm.Open(dialector)
	if err != nil {
		log.Fatal(err)
	}

	return dbMock, sqlMock
}

func TestGetAllUsersRepository(t *testing.T) {
	tests := []userRepositoryTest{
		{
			description:   "Should return no error on get all users",
			expectedQuery: "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL",
		},
		{
			description:   "Should return error on get all users",
			expectedQuery: "SELECT * FROM `users` WHERE`users`.`deleted_at` IS NULL",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Assert
			mockUser := entity.User{}
			faker.Struct(&mockUser)

			rows := sqlmock.NewRows([]string{
				"ID",
				"FirstName",
				"LastName",
				"Email",
				"Password",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockUser.ID,
				mockUser.FirstName,
				mockUser.LastName,
				mockUser.Email,
				mockUser.Password,
				mockUser.CreatedAt,
				mockUser.UpdatedAt,
				nil,
			)

			query := regexp.QuoteMeta(testCase.expectedQuery)
			dbMock, sqlMock := mockDb()

			sqlMock.ExpectQuery(query).WillReturnRows(rows)

			// Act
			userRepository := persistence.NewUserRepository(dbMock)
			_, err := userRepository.GetAllUsers()
			// Assert
			if err != nil {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
