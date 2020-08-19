package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// IsAuthenticated middleware that verify the user if he's authenticated
func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorObject := &models.RequestError{}
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
				errorObject.StatusCode = http.StatusUnauthorized
				utils.RespondWithError(w, errorObject)
				return
			}

			if token.Valid {
				claims, _ := token.Claims.(jwt.MapClaims)
				ctx := context.WithValue(r.Context(), userCtxKey, claims["id"].(string))
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = "Invalid token"
				errorObject.StatusCode = http.StatusUnauthorized
				utils.RespondWithError(w, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token."
			errorObject.StatusCode = http.StatusUnauthorized
			utils.RespondWithError(w, errorObject)
			return
		}
	})
}

// GetUserID get the user id from the context
func GetUserID(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
