package routes

import (
	"net/http"

	"github.com/Eugene600/Go-Project/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() http.Handler {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	//AUTH
	router.POST("/auth/signup", handlers.SignUpUser)

	// USERS
	router.GET("/users/search", handlers.GetUserByUsername)

	router.GET("/users", handlers.GetAllUsers)

	router.PUT("/users/:id", handlers.UpdateUser)

	router.DELETE("/users/:id", handlers.DeleteUser)

	router.PUT("/users/:id/recover", handlers.RecoverDeletedUser)

	router.GET("/users/deleted", handlers.GetDeletedUsers)

	return router
}
