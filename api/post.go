package api

import (
	"encoding/json"
	"net/http"

	"github.com/lazhari/web-jwt/middleware"
	"github.com/lazhari/web-jwt/models"
	"github.com/lazhari/web-jwt/post"
	"github.com/lazhari/web-jwt/utils"
)

// PostHandler the post hander interface
type PostHandler interface {
	CreatePost(http.ResponseWriter, *http.Request)
	GetAllPosts(http.ResponseWriter, *http.Request)
}

type postHandler struct {
	postService post.Service
}

// NewPostHandler creates a new post handler
func NewPostHandler(postSrv post.Service) PostHandler {
	return &postHandler{postService: postSrv}
}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.GetUserID(ctx)
	post := &models.Post{}

	json.NewDecoder(r.Body).Decode(post)
	post.AuthorID = userID

	post, err := p.postService.Create(post)

	if err != nil {
		utils.RespondWithError(w, err)
	}

	utils.ResponseJSON(w, post)
	return
}

func (p *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {

}
