package handlers

import (
	"finance-backend/config"
	"finance-backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateUserRole(c *gin.Context) {
	log.Println("User: UpdateUserRole started")

	id := c.Param("id")

	var body struct {
		Role string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("User: updating role id=%s new_role=%s", id, body.Role)

	_, err := config.DB.Exec(
		"UPDATE users SET role=$1 WHERE id=$2",
		body.Role,
		id,
	)

	if err != nil {
		log.Printf("User: update failed id=%s error=%v", id, err)

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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
