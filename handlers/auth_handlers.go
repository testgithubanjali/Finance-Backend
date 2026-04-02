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

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Signup: invalid request payload", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	log.Printf("Signup: received signup request for email=%s", user.Email)

	err := services.Signup(user)
	if err != nil {
		log.Printf("Signup: failed to create user email=%s: %v", user.Email, err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Signup: user created successfully email=%s", user.Email)
	c.JSON(201, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	log.Println("Login starts")
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Login: invalid request payload", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	log.Printf("Login: request for email=%s", req.Email)

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		log.Printf("Login: invalid credentials for email=%s: %v", req.Email, err)
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Printf("Login: successful for email=%s", req.Email)
	c.JSON(200, gin.H{"token": token})
}
