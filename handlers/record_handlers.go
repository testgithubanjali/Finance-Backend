package handlers

import (
	"finance-backend/config"
	"finance-backend/models"
	"finance-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user.ID = uuid.New().String()
	user.Password = utils.HashPassword(user.Password)
	user.Role = "viewer"

	query := `
	INSERT INTO users (id, name, email, password, role)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := config.DB.Exec(query, user.ID, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "User creation failed"})
		return
	}

	c.JSON(201, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var req models.User

	c.ShouldBindJSON(&req)

	var user models.User

	query := `SELECT id, password, role FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.Password, &user.Role)

	if err != nil || !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := utils.GenerateToken(user.ID, user.Role)

	c.JSON(200, gin.H{"token": token})
}
