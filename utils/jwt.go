package utils

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("secret")

func GenerateToken(userID, role string) string {
	log.Printf("JWT: GenerateToken started user_id=%s role=%s", userID, role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	})

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		log.Printf("JWT: GenerateToken failed user_id=%s error=%v", userID, err)
		return ""
	}

	log.Printf("JWT: GenerateToken succeeded user_id=%s", userID)
	return tokenStr
}

func ValidateToken(tokenStr string) (string, string, error) {
	log.Println("JWT: ValidateToken started")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		log.Printf("JWT: ValidateToken failed parse error=%v", err)
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("JWT: ValidateToken failed invalid claims")
		return "", "", jwt.ErrInvalidKey
	}

	userID, ok1 := claims["user_id"].(string)
	role, ok2 := claims["role"].(string)
	if !ok1 || !ok2 {
		log.Println("JWT: ValidateToken failed missing claims")
		return "", "", jwt.ErrTokenMalformed
	}

	log.Printf("JWT: ValidateToken succeeded user_id=%s role=%s", userID, role)
	return userID, role, nil
}
