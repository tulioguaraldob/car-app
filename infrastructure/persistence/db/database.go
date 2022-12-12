package db

import (
	"fmt"
	"log"

	"github.com/TulioGuaraldoB/car-app/config/env"
	"github.com/TulioGuaraldoB/car-app/domain/repository"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence"
	"github.com/TulioGuaraldoB/car-app/infrastructure/persistence/db/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	db              *gorm.DB
	IUserRepository repository.IUserRepository
	ICarRepository  repository.ICarRepository
}

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var db *gorm.DB

func NewRepositories() *Repositories {
	OpenConnection()

	return &Repositories{
		db:              db,
		IUserRepository: persistence.NewUserRepository(db),
		ICarRepository:  persistence.NewCarRepository(db),
	}
}

func StartDataBase(dbConfig DbConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect to MySql Server. %s", err.Error())
		log.Fatal(errMessage)
	}

	db = database

	migration.Run(db)
}

func OpenConnection() *gorm.DB {
	if db == nil {
		StartDataBase(DbConfig{
			User:     env.Env.DbUser,
			Password: env.Env.DbPassword,
			Host:     env.Env.DbHost,
			Port:     env.Env.DbPort,
			Name:     env.Env.DbName,
		})
	}

	return db
}
