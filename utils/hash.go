package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	log.Println("HashUtil: HashPassword started")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Printf("HashUtil: HashPassword failed: %v", err)
		return ""
	}
	log.Println("HashUtil: HashPassword succeeded")
	return string(hash)
}

func CheckPassword(hash, password string) bool {
	log.Println("HashUtil: CheckPassword started")
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Printf("HashUtil: CheckPassword failed: %v", err)
		return false
	}
	log.Println("HashUtil: CheckPassword succeeded")
	return true
}
