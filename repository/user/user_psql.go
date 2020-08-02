package user

import (
	"database/sql"

	"github.com/lazhari/web-jwt/models"
)

// UserRepository The user repository
type UserRepository struct{}

// SignUp Create a user row in the users table
func (u UserRepository) SignUp(db *sql.DB, user models.User) (models.User, error) {
	stmt := "INSERT INTO users (email, password) values($1, $2) RETURNING id;"

	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

// Login Get the user from the users table
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	return user, err
}
