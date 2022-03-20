package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madxiii/mongocrud/internal/models"
)

func (s *Server) GetUsers(c *gin.Context) {
	users, err := s.Store.Find(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": true,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": users,
	})
}

func (s *Server) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	if err := s.Store.Insert(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "user created",
	})
}

func (s *Server) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if err := s.Store.Update(c.Request.Context(), id); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "user updated",
	})
}

func (s *Server) DeleteUser(c *gin.Context) {
	fmt.Println("Test")
}
