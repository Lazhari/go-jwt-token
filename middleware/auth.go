package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
)

// IsAuthenticated middleware that verify the user if he's authenticated
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorObject := models.Error{}
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if error != nil {
				errorObject.Message = "Invalid token"
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = "Invalid token"
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
