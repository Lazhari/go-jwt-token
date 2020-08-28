package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazhari/web-jwt/utils"

	"github.com/lazhari/web-jwt/models"

	"github.com/jinzhu/gorm"
	"github.com/lazhari/web-jwt/middleware"
)

// CreatePost creates a new post
func (c Controller) CreatePost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID := middleware.GetUserID(ctx)
		post := models.Post{}
		err := &models.RequestError{}

		json.NewDecoder(r.Body).Decode(&post)
		post.AuthorID = userID

		errV := post.Validate()

		if errV != nil {
			utils.RespondWithError(w, &models.RequestError{
				StatusCode:       http.StatusBadRequest,
				Message:          "Invalid request",
				ValidationErrors: errV,
			})
			return
		}

		dbc := db.Create(&post)

		if dbc.Error != nil {
			err.Message = "Internal server error"
			err.StatusCode = http.StatusInternalServerError

			utils.RespondWithError(w, err)
		}

		utils.ResponseJSON(w, post)
		return
	}
}
