package user

import "github.com/lazhari/web-jwt/models"

// Repository the auth repository
type Repository interface {
	SignUp(models.User) (models.User, error)
	Login(models.User) (models.User, error)
}
