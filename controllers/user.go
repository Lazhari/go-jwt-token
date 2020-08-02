package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	userRepository "github.com/lazhari/web-jwt/repository/user"

	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
	"golang.org/x/crypto/bcrypt"
)

// SignUpHandler The user sign up handler
func (c Controller) SignUpHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := models.RequestError{}
		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			err.Message = "Email is missing."
			err.StatusCode = http.StatusBadRequest
			utils.RespondWithError(w, err)
			return
		}

		if !utils.IsEmailValid(user.Email) {
			err.Message = "Email is not valid."
			err.StatusCode = http.StatusBadRequest
			utils.RespondWithError(w, err)
			return
		}

		if user.Password == "" {
			err.Message = "Password is missing."
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

		userRepo := userRepository.UserRepository{}
		user, errInsert := userRepo.SignUp(db, user)

		if errInsert != nil {
			log.Printf("Error while inserting the user into db: %v\n", errInsert)
			err.Message = "Internal server error"
			err.StatusCode = http.StatusInternalServerError
			utils.RespondWithError(w, err)
			return
		}

		utils.ResponseJSON(w, user)
	}
}

// LoginHandler The login handler
func (c Controller) LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		jwt := models.JWT{}
		error := models.RequestError{}

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

		userRepo := userRepository.UserRepository{}

		user, err := userRepo.Login(db, user)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				error.StatusCode = http.StatusNotFound
			} else {
				error.Message = err.Error()
				error.StatusCode = http.StatusInternalServerError
			}
			utils.RespondWithError(w, error)
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
}
