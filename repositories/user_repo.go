package repositories

import (
	"finance-backend/config"
	"finance-backend/models"
)

func CreateUser(user models.User) error {
	query := `INSERT INTO users (id, name, email, password, role)
	          VALUES ($1,$2,$3,$4,$5)`

	_, err := config.DB.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Role,
	)

	if err != nil {
		return err // 🔥 MUST return error
	}

	return nil
}
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT id, password, role FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Password, &user.Role)

	return user, err
}
