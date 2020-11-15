package user

import "github.com/lazhari/web-jwt/models"

// Service the authentication service
type Service interface {
	SignUp(models.User) (models.User, error)
	Login(models.User) (models.User, error)
}
