package middleware

import (
	"finance-backend/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			log.Println("AuthMiddleware: missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			log.Printf("AuthMiddleware: invalid Authorization header format: %q", authHeader)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token := parts[1]

		userID, role, err := utils.ValidateToken(token)
		if err != nil {
			log.Printf("AuthMiddleware: token validation failed: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Printf("AuthMiddleware: authenticated user_id=%s role=%s", userID, role)
		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}
