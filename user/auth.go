package user

import "github.com/lazhari/web-jwt/models"

type userService struct {
	authRepo Repository
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo Repository) Service {
	return &userService{
		userRepo,
	}
}

func (authSrv *userService) Login(user models.User) (models.User, error) {
	return authSrv.authRepo.Login(user)
}

func (authSrv *userService) SignUp(user models.User) (models.User, error) {
	return authSrv.authRepo.SignUp(user)
}
