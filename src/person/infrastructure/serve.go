package infrastructure

import (
	"api/src/person/application"
	"api/src/person/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	repo, err := NewPostgresPersonRepository()
	if err != nil {
		log.Fatalf("Error al conectar DB: %v", err)
	}

    addPersonService := application.NewAddPersonService(repo)
	listPersonService := application.NewListPersonService(repo)
	statsService := application.NewStatsService(repo)

	routes.SetupRoutes(router, addPersonService, listPersonService, statsService)

	return &Server{
		router: router,
	}
}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}
