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
func RespondWithError(w http.ResponseWriter, err error) {
	switch error := err.(type) {
	case *models.RequestError:
		w.WriteHeader(error.StatusCode)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
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
		"id":    user.ID,
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

// ParseToken get the user id from the token
func ParseToken(tokenStr string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := claims["id"].(string)
		return id, nil
	}

	return "", err
}
