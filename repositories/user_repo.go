package repositories

import (
	"finance-backend/config"
	"finance-backend/models"
	"log"
)

func CreateUser(user models.User) error {
	log.Printf("UserRepo: CreateUser email=%s role=%s", user.Email, user.Role)
	query := `INSERT INTO users (id, name, email, password, role)
	          VALUES ($1,$2,$3,$4,$5)`

	_, err := config.DB.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Role,
	)

	if err != nil {
		log.Printf("UserRepo: CreateUser failed email=%s error=%v", user.Email, err)
		return err
	}

	log.Printf("UserRepo: CreateUser succeeded email=%s id=%s", user.Email, user.ID)
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT id, password, role FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Password, &user.Role)
	if err != nil {
		log.Printf("UserRepo: GetUserByEmail failed email=%s error=%v", email, err)
		return user, err
	}

	log.Printf("UserRepo: GetUserByEmail found id=%s email=%s role=%s", user.ID, email, user.Role)
	return user, nil
}
func UpdateUserStatus(id string, isActive bool) error {
	_, err := config.DB.Exec(
		"UPDATE users SET is_active=$1 WHERE id=$2",
		isActive, id,
	)
	return err
}
