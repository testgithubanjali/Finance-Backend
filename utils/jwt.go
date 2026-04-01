package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("secret")

func GenerateToken(userID, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	})

	tokenString, _ := token.SignedString(SECRET)
	return tokenString
}

func ValidateToken(tokenStr string) (string, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return "", "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["user_id"].(string), claims["role"].(string), nil
}
