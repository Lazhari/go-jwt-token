package user

import (
	"log"
	"net/http"

	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	authRepo Repository
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo Repository) Service {
	return &userService{
		userRepo,
	}
}

func (authSrv *userService) Login(user *models.User) (*models.User, error) {
	return authSrv.authRepo.Login(user)
}

func (authSrv *userService) SignUp(user *models.User) (*models.User, error) {
	err := &models.RequestError{}
	Validationerr := user.Validate()
	if Validationerr != nil {
		return nil, Validationerr
	}

	if !utils.IsEmailValid(user.Email) {
		err.Message = "Email is not valid."
		err.StatusCode = http.StatusBadRequest
		return nil, err
	}

	hash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if errHash != nil {
		log.Printf("Error while hashing the password: %v\n", errHash)
		err.Message = "Internal Server Error"
		err.StatusCode = http.StatusInternalServerError
		return nil, err
	}

	user.Password = string(hash)
	return authSrv.authRepo.SignUp(user)
}
