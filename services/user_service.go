package services

import (
	"finance-backend/models"
	"finance-backend/repositories"
	"finance-backend/utils"
	"log"

	"github.com/google/uuid"
)

func Signup(user models.User) error {
	log.Println("Signup: started")

	user.ID = uuid.New().String()
	log.Printf("Signup: generated user ID = %s", user.ID)

	user.Password = utils.HashPassword(user.Password)
	log.Println("Signup: password hashed")

	user.Role = "viewer"
	log.Println("Signup: assigned default role = viewer")

	err := repositories.CreateUser(user)
	if err != nil {
		log.Printf("Signup: failed to create user, error = %v", err)
		return err
	}

	log.Println("Signup: user created successfully")
	return nil
}

func Login(email, password string) (string, error) {
	log.Printf("Login: attempt for email = %s", email)

	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		log.Printf("Login: user not found or DB error = %v", err)
		return "", err
	}

	log.Printf("Login: user found, ID = %s", user.ID)

	if !utils.CheckPassword(user.Password, password) {
		log.Println("Login: password mismatch")
		return "", err
	}

	log.Println("Login: password verified")

	token := utils.GenerateToken(user.ID, user.Role)
	log.Println("Login: token generated")

	return token, nil
}
