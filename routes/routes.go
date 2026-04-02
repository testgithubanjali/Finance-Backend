package routes

import (
	"finance-backend/handlers"
	"finance-backend/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	log.Println("Routes: setting up HTTP routes")

	r.GET("/health", func(c *gin.Context) {
		log.Println("Routes: /health hit")
		c.JSON(200, gin.H{"status": "OK"})
	})

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	{
		auth.GET("/records", handlers.GetRecords)
		auth.POST("/records", middleware.Authorize("admin"), handlers.CreateRecord)

		auth.GET("/dashboard", handlers.GetSummary)
		auth.PUT("/records/:id", middleware.Authorize("admin"), handlers.UpdateRecord)
		auth.DELETE("/records/:id", middleware.Authorize("admin"), handlers.DeleteRecord)
		auth.PATCH("/users/:id/role", middleware.Authorize("admin"), handlers.UpdateUserRole)
		auth.PATCH("/users/:id/status", middleware.Authorize("admin"), handlers.UpdateUserStatus)
	}

	log.Println("Routes: route setup complete")
}
