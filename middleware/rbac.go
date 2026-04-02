package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		log.Printf("RBAC: checking role=%v allowed=%v", role, roles)
		for _, r := range roles {
			if r == role {
				log.Printf("RBAC: authorized role=%v", role)
				c.Next()
				return
			}
		}

		log.Printf("RBAC: forbidden role=%v allowed=%v", role, roles)
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
