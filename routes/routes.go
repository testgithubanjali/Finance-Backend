package routes

import (
	"finance-backend/handlers"
	"finance-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	{
		auth.GET("/records", handlers.GetRecords)
		auth.POST("/records", middleware.Authorize("admin"), handlers.CreateRecord)

		auth.GET("/dashboard", handlers.GetSummary)
	}
}
