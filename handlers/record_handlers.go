package handlers

import (
	"finance-backend/models"
	"finance-backend/services"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
	var record models.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	userID, _ := c.Get("user_id")

	err := services.CreateRecord(record, userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed"})
		return
	}

	c.JSON(201, gin.H{"message": "Created"})
}

func GetRecords(c *gin.Context) {
	records, _ := services.GetRecords()
	c.JSON(200, records)
}
