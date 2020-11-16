package post

import "github.com/lazhari/web-jwt/models"

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
	return postSrv.postRepo.Create(p)
}

func (postSrv *postService) GetAll() ([]models.Post, error) {
	return postSrv.postRepo.GetAll()
}
