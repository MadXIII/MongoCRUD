package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/mongocrud/internal/database"
)

type Server struct {
	Store  database.Store
	Router *gin.Engine
}

func Constr(store database.Store) *Server {
	return &Server{store, gin.Default()}
}

func (s *Server) Routes() {
	group := s.Router.Group("/users")
	{
		group.GET("/", s.GetUsers)
		group.POST("/", s.CreateUser)
		group.PUT("/:id", s.UpdateUser)
		group.DELETE("/:id", s.DeleteUser)
	}

	if err := s.Router.Run(":8989"); err != nil {
		log.Fatal(err.Error())
	}
}
