package userrepo

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/lazhari/web-jwt/driver"
	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/user"
)

// Repository The user repository
type Repository struct{}

type postgresRepository struct {
	db *gorm.DB
}

// NewPostgreRepository Create the postgres repository
func NewPostgreRepository() (user.Repository, error) {
	db, err := driver.ConnectDB()
	if err != nil {
		return nil, err
	}

	repo := postgresRepository{
		db,
	}
	return repo, nil
}

// SignUp Create a user row in the users table
func (pr postgresRepository) SignUp(user *models.User) (*models.User, error) {
	// stmt := "INSERT INTO users (email, password) values($1, $2) RETURNING id;"

	// err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	results := int64(0)
	pr.db.Model(&models.User{}).Where("email = ?", user.Email).Count(&results)

	if results > 0 {
		return user, &models.RequestError{
			Message:    fmt.Sprintf("This email %q already exist", user.Email),
			StatusCode: http.StatusBadRequest,
		}
	}

	dbc := pr.db.Create(&user)

	if dbc.Error != nil {
		return user, &models.RequestError{
			Message:    "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}
	}

	user.Password = ""
	return user, nil
}

// Login Get the user from the users table
func (pr postgresRepository) Login(user *models.User) (*models.User, error) {
	// row := db.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)
	// err := row.Scan(&user.ID, &user.Email, &user.Password)

	dbc := pr.db.First(&user, "email = ?", user.Email)

	if dbc.Error != nil {
		errorReq := &models.RequestError{}
		if dbc.Error == sql.ErrNoRows {
			errorReq.Message = "The user does not exist!"
			errorReq.StatusCode = http.StatusNotFound
		} else {
			errorReq.Message = dbc.Error.Error()
			errorReq.StatusCode = http.StatusInternalServerError
		}
		return user, errorReq
	}

	return user, nil
}
