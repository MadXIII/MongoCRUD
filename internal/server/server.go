package server

import (
	"github.com/gin-gonic/gin"
	"github.com/madxiii/mongocrud/internal/database"
	"github.com/madxiii/mongocrud/internal/handlers"
)

type Server struct {
	Store  database.Store
	Router *gin.Engine
}

func Constr(store database.Store) *Server {
	return &Server{store, gin.Default()}
}

func (s *Server) Routes() {
	s.Router.Group("/users")
	{
		s.Router.GET("/:id", handlers.CreateUser)
		s.Router.POST("/", handlers.CreateUser)
		s.Router.PUT("/", handlers.CreateUser)
		s.Router.DELETE("/", handlers.CreateUser)
	}
}
