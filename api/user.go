package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazhari/web-jwt/user"

	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/utils"
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
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	user, err := h.authService.SignUp(user)

	if err != nil {
		log.Printf("Error while inserting the user into db: %v\n", err)
		utils.RespondWithError(w, err)
		return
	}

	utils.ResponseJSON(w, user)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	json.NewDecoder(r.Body).Decode(user)

	jwt, err := h.authService.Login(user)

	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	utils.ResponseJSON(w, jwt)
}
