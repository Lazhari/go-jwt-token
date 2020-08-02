package user

import (
	"database/sql"
	"net/http"

	"github.com/lazhari/web-jwt/models"
)

// Repository The user repository
type Repository struct{}

// SignUp Create a user row in the users table
func (u Repository) SignUp(db *sql.DB, user models.User) (models.User, error) {
	stmt := "INSERT INTO users (email, password) values($1, $2) RETURNING id;"

	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

// Login Get the user from the users table
func (u Repository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		errorReq := &models.RequestError{}
		if err == sql.ErrNoRows {
			errorReq.Message = "The user does not exist!"
			errorReq.StatusCode = http.StatusNotFound
		} else {
			errorReq.Message = err.Error()
			errorReq.StatusCode = http.StatusInternalServerError
		}
		return user, errorReq
	}

	return user, nil
}
