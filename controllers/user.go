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
		err := models.Error{}
		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			err.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		if !utils.IsEmailValid(user.Email) {
			err.Message = "Email is not valid."
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		if user.Password == "" {
			err.Message = "Password is missing."
			utils.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		hash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if errHash != nil {
			log.Printf("Error while hashing the password: %v\n", errHash)
			err.Message = "Internal Server Error"
			utils.RespondWithError(w, http.StatusInternalServerError, err)
		}
		user.Password = string(hash)

		userRepo := userRepository.UserRepository{}
		user, errInsert := userRepo.SignUp(db, user)

		if errInsert != nil {
			log.Printf("Error while inserting the user into db: %v\n", errInsert)
			err.Message = "Internal server error"
			utils.RespondWithError(w, http.StatusInternalServerError, err)
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
		error := models.Error{}

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing."
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "Password is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		password := user.Password

		userRepo := userRepository.UserRepository{}

		user, err := userRepo.Login(db, user)

		if err != nil {
			var status int
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				status = http.StatusNotFound
			} else {
				error.Message = err.Error()
				status = http.StatusInternalServerError
			}
			utils.RespondWithError(w, status, error)
			return
		}

		hashedPassword := user.Password

		ok := utils.ComparePasswords(hashedPassword, []byte(password))

		if !ok {
			error.Message = "The password isn't valid"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}
		jwt.Token = token
		utils.ResponseJSON(w, jwt)
	}
}
