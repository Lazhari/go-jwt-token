package post

import (
	"net/http"

	"github.com/lazhari/web-jwt/models"
)

type postService struct {
	postRepo Repository
}

// NewPostService create a new post service
func NewPostService(postRepo Repository) Service {
	return &postService{
		postRepo,
	}
}

func (postSrv *postService) Create(p *models.Post) (*models.Post, error) {
	err := &models.RequestError{}

	validationErr := p.Validate()

	if validationErr != nil {
		return nil, &models.RequestError{
			StatusCode:       http.StatusBadRequest,
			Message:          "Invalid request",
			ValidationErrors: validationErr,
		}
	}

	p, createError := postSrv.postRepo.CreatePost(p)
	if createError != nil {
		err.Message = "Internal server error"
		err.StatusCode = http.StatusInternalServerError

		return nil, err
	}

	return p, nil
}

func (postSrv *postService) GetAll() ([]models.Post, error) {
	return postSrv.postRepo.GetAllPosts()
}
