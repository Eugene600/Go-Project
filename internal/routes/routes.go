package routes

import (
	"net/http"

	"github.com/Eugene600/Go-Project/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() http.Handler {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	// USERS
	router.POST("/users", handlers.CreateUser)

	router.GET("/users/search", handlers.GetUserByUsername)

	router.GET("/users", handlers.GetAllUsers)

	return router
}
