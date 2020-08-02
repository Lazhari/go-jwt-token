package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/lazhari/web-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

// RespondWithError Response with an error
func RespondWithError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// ResponseJSON returns a json response
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func ComparePasswords(hashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
