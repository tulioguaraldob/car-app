package main

import (
	"github.com/TulioGuaraldoB/car-app/config/env"
	"github.com/TulioGuaraldoB/car-app/infrastructure/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.New()
	server.Run()
}
