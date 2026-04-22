package repositories

import (
	config "taskmanager/config"
	"taskmanager/models"
)

func CreateUser(Email string, password string) error {
	query := "INSERT INTO users(email,password) VALUES ($1,$2)"
	_, err := config.DB.Exec(query, Email, password)
	return err
}
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := config.DB.QueryRow(query, email).
		Scan(&user.Id, &user.Email, &user.Password)

	return user, err
}
