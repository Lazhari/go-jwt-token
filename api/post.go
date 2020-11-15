package api

// CreatePost creates a new post
/* func (c Controller) CreatePost(db *gorm.DB) http.HandlerFunc {
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
} */

// GetPostByID retrieve a post by ID
/* func (c Controller) GetPostByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := &models.RequestError{}
		vars := mux.Vars(r)

		id := vars["id"]
		if id == "" {
			err.Message = "ID is required"
			err.StatusCode = http.StatusBadRequest

			utils.RespondWithError(w, err)
		}

		post := models.Post{}
		dbc := db.First(&post, "id = ?", id)

		if dbc.Error != nil {
			if dbc.Error == sql.ErrNoRows {
				err.Message = "The user does not exist!"
				err.StatusCode = http.StatusNotFound
			} else {
				err.Message = dbc.Error.Error()
				err.StatusCode = http.StatusInternalServerError
			}
			utils.RespondWithError(w, err)
			return
		}

		utils.ResponseJSON(w, post)
		return
	}
} */
