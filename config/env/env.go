package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVar struct {
	DbUser        string
	DbPassword    string
	DbHost        string
	DbPort        string
	DbName        string
	Port          string
	CrossroadsUrl string
}

var Env *EnvironmentVar

func GetEnvironmentVariables() *EnvironmentVar {
	godotenv.Load(".env")

	Env = &EnvironmentVar{
		DbUser:        os.Getenv("DB_USER"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbHost:        os.Getenv("DB_HOST"),
		DbPort:        os.Getenv("DB_PORT"),
		DbName:        os.Getenv("DB_NAME"),
		Port:          os.Getenv("PORT"),
		CrossroadsUrl: os.Getenv("CROSS_ROADS_URL"),
	}

	return Env
}
