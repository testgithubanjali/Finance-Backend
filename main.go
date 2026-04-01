package main

import (
	"finance-backend/config"
	"finance-backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Starting Finance Backend...")

	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("🌐 Server running on port 8080")
	r.Run(":8080")
}
