package routes

import (
	"net/http"

	"github.com/Eugene600/Go-Project/internal/handlers"
	"github.com/Eugene600/Go-Project/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetRoutes() http.Handler {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	//AUTH
	auth := router.Group("/auth")

	auth.POST("/signup", handlers.SignUpUser)

	auth.POST("/login", handlers.LoginUser)

	// USERS
	users := router.Group("/users")
	users.Use(middlewares.AuthMiddleware())

	users.GET("/search", handlers.GetUserByUsername)

	users.GET("", handlers.GetAllUsers)

	users.PUT("/:id", handlers.UpdateUser)

	users.DELETE("/:id", handlers.DeleteUser)

	users.PUT("/:id/recover", handlers.RecoverDeletedUser)

	users.GET("/deleted", handlers.GetDeletedUsers)

	return router
}
