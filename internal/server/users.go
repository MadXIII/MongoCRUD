package server

import (
	"github.com/gin-gonic/gin"
	"github.com/madxiii/mongocrud/internal/models"
)

func (s *Server) GetUsers(c *gin.Context) {
	users, err := s.Store.Find(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
		})
	}
	c.JSON(200, users)
}

func (s *Server) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
	}
	s.Store.Insert(c.Request.Context(), user)

	c.JSON(200, gin.H{
		"error": false,
	})
}

func (s *Server) UpdateUser(c *gin.Context) {
}

func (s *Server) DeleteUser(c *gin.Context) {
}
