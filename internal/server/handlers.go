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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}

func (s *Server) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	if err := s.Store.Insert(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func (s *Server) UpdateUser(c *gin.Context) {
	var newUser models.User
	id := c.Param("id")
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}
	if err := s.Store.Update(c.Request.Context(), id, newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

func (s *Server) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := s.Store.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Deleted",
	})
}
