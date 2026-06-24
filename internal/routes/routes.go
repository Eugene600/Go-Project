package routes

import (
	"net/http"

	"github.com/Eugene600/Go-Project/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() http.Handler {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// USERS
	router.POST("/users", handlers.CreateUser)

	return router
}
