package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazhari/web-jwt/user"

	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler the user handler interface
type UserHandler interface {
	SignUp(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type handler struct {
	authService user.Service
}

// NewHandler Creates a new http handler
func NewHandler(authSrv user.Service) UserHandler {
	return &handler{authService: authSrv}
}

// SignUpHandler The user sign up handler
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := &models.RequestError{}
	json.NewDecoder(r.Body).Decode(&user)

	errV := user.Validate()

	if errV != nil {
		utils.RespondWithError(w, &models.RequestError{
			StatusCode:       http.StatusBadRequest,
			Message:          "Invalid request",
			ValidationErrors: errV,
		})
		return
	}

	if !utils.IsEmailValid(user.Email) {
		err.Message = "Email is not valid."
		err.StatusCode = http.StatusBadRequest
		utils.RespondWithError(w, err)
		return
	}

	hash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if errHash != nil {
		log.Printf("Error while hashing the password: %v\n", errHash)
		err.Message = "Internal Server Error"
		err.StatusCode = http.StatusInternalServerError
		utils.RespondWithError(w, err)
	}
	user.Password = string(hash)

	user, errInsert := h.authService.SignUp(user)

	if errInsert != nil {
		log.Printf("Error while inserting the user into db: %v\n", errInsert)
		utils.RespondWithError(w, errInsert)
		return
	}

	utils.ResponseJSON(w, user)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	jwt := models.JWT{}
	error := &models.RequestError{}

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		error.StatusCode = http.StatusBadRequest
		utils.RespondWithError(w, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		error.StatusCode = http.StatusBadRequest
		utils.RespondWithError(w, error)
		return
	}

	password := user.Password

	user, err := h.authService.Login(user)

	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	hashedPassword := user.Password

	ok := utils.ComparePasswords(hashedPassword, []byte(password))

	if !ok {
		error.Message = "The password isn't valid"
		error.StatusCode = http.StatusUnauthorized
		utils.RespondWithError(w, error)
		return
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		error.Message = err.Error()
		error.StatusCode = http.StatusInternalServerError
		utils.RespondWithError(w, error)
		return
	}
	jwt.Token = token
	utils.ResponseJSON(w, jwt)
}
