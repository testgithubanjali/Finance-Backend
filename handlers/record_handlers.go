package handlers

import (
	"finance-backend/models"
	"finance-backend/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
	log.Println("Record: CreateRecord started")

	var record models.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		log.Printf("Record: invalid payload: %v", err)
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	uid, ok := userIDVal.(string)
	if !ok || uid == "" {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}

	log.Printf("Record: CreateRecord user_id=%s type=%s amount=%.2f", uid, record.Type, record.Amount)

	err := services.CreateRecord(record, uid)
	if err != nil {
		log.Printf("Record: failed user_id=%s error=%v", uid, err)
		c.JSON(500, gin.H{"error": "Failed"})
		return
	}

	log.Printf("Record: succeeded user_id=%s id=%s", uid, record.ID)
	c.JSON(201, gin.H{"message": "Created"})
}
func GetRecords(c *gin.Context) {
	log.Println("Record: GetRecords started")

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(string)

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	typeFilter := c.Query("type")
	category := c.Query("category")

	var records []models.Record
	var err error

	records, err = services.GetFilteredRecords(userID, typeFilter, category, limit, offset)

	if err != nil {
		log.Printf("Record: GetRecords failed: %v", err)
		c.JSON(500, gin.H{"error": "Failed to fetch records"})
		return
	}

	log.Printf("Record: returned %d records", len(records))

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  records,
	})
}
func UpdateRecord(c *gin.Context) {
	id := c.Param("id")

	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err := services.UpdateRecord(id, record)
	if err != nil {
		log.Printf("Update error: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Updated"})
}
func DeleteRecord(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteRecord(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted"})
}
