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

func (authSrv *userService) Login(user *models.User) (*models.JWT, error) {
	resErr := &models.RequestError{}
	jwt := models.JWT{}
	if user.Email == "" {
		resErr.Message = "Email is missing."
		resErr.StatusCode = http.StatusBadRequest
		return nil, resErr
	}

	if user.Password == "" {
		resErr.Message = "Password is missing"
		resErr.StatusCode = http.StatusBadRequest
		return nil, resErr
	}

	password := user.Password
	user, err := authSrv.authRepo.Login(user)

	if err != nil {
		return nil, err
	}

	hashedPassword := user.Password
	ok := utils.ComparePasswords(hashedPassword, []byte(password))

	if !ok {
		resErr.Message = "Invalid credentials!"
		resErr.StatusCode = http.StatusUnauthorized
		return nil, resErr
	}

	token, err := utils.GenerateToken(*user)

	if err != nil {
		resErr.Message = err.Error()
		resErr.StatusCode = http.StatusInternalServerError
		return nil, resErr
	}
	jwt.Token = token
	return &jwt, nil
}

func (authSrv *userService) SignUp(user *models.User) (*models.User, error) {
	err := &models.RequestError{}
	validationErr := user.Validate()
	if validationErr != nil {
		return nil, &models.RequestError{
			StatusCode:       http.StatusBadRequest,
			Message:          "Invalid request",
			ValidationErrors: validationErr,
		}
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
