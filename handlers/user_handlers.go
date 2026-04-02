package handlers

import (
	"finance-backend/config"
	"finance-backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateUserRole(c *gin.Context) {
	id := c.Param("id")
	log.Printf("UserHandler: UpdateUserRole started id=%s", id)

	var body struct {
		Role string `json:"role"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("UserHandler: UpdateUserRole invalid input id=%s error=%v", id, err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("UserHandler: UpdateUserRole request id=%s role=%s", id, body.Role)
	_, err := config.DB.Exec(
		"UPDATE users SET role=$1 WHERE id=$2",
		body.Role, id,
	)

	if err != nil {
		log.Printf("UserHandler: UpdateUserRole failed id=%s role=%s error=%v", id, body.Role, err)
		c.JSON(500, gin.H{"error": "Failed"})
		return
	}

	log.Printf("UserHandler: UpdateUserRole succeeded id=%s role=%s", id, body.Role)
	c.JSON(200, gin.H{"message": "Role updated"})
}

func UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")
	log.Printf("UserHandler: UpdateUserStatus started id=%s", id)

	var body struct {
		IsActive bool `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("UserHandler: UpdateUserStatus invalid input id=%s error=%v", id, err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("UserHandler: UpdateUserStatus request id=%s active=%v", id, body.IsActive)
	err := services.UpdateUserStatus(id, body.IsActive)
	if err != nil {
		log.Printf("UserHandler: UpdateUserStatus failed id=%s active=%v error=%v", id, body.IsActive, err)
		c.JSON(500, gin.H{"error": "Failed to update status"})
		return
	}

	log.Printf("UserHandler: UpdateUserStatus succeeded id=%s active=%v", id, body.IsActive)
	c.JSON(200, gin.H{"message": "User status updated"})
}
