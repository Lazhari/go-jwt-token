package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lazhari/web-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

// RespondWithError Response with an error
func RespondWithError(w http.ResponseWriter, err models.RequestError) {
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}

// ResponseJSON returns a json response
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

// ComparePasswords Check the user password
func ComparePasswords(hashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GenerateToken Generate a new toke
func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
		"sub":   "user",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 40).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", &models.RequestError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server Error",
		}
	}
	return tokenString, nil
}
