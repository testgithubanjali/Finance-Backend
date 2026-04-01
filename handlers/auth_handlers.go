package handlers

import (
	"finance-backend/models"
	"finance-backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	log.Println("Signup starts")
	var user models.User
	log.Println("Signup: Signup request: ")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.Signup(user)
	if err != nil {
		log.Println("Signup: failed to create user", err)
		c.JSON(500, gin.H{"error": "Failed"})
		return
	}
	log.Println("Signup: user created successfully")

	c.JSON(201, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	log.Println("Login starts")
	var req models.User
	c.ShouldBindJSON(&req)
	log.Println("Login: request for email and password: ")

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		log.Println("Login: invalid credentials", err)
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Println("Login: successful")
	c.JSON(200, gin.H{"token": token})
}
