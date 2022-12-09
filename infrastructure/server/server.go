package server

import (
	"log"

	"github.com/TulioGuaraldoB/car-app/config/env"
	"github.com/TulioGuaraldoB/car-app/infrastructure/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func New() Server {
	return Server{
		Port:   env.Env.Port,
		Server: routes.GetRoutes(),
	}
}

func (s *Server) Run() {
	log.Fatal(s.Server.Run(":" + s.Port))
}
