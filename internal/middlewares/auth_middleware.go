package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Eugene600/Go-Project/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			log.Print("Missing authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Print("Missing Bearer prefix")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.VerifyJwt(tokenString)

		if err != nil {
			log.Printf("Error from validating token: %s", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid access token",
			})
			return
		}

		userId, err := utils.GetUserIdFromToken(token)
		if err != nil {
			log.Printf("Error while obtaining user id from token, %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
