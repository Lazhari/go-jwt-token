package post

import "github.com/lazhari/web-jwt/models"

// Repository the post repository
type Repository interface {
	CreatePost(*models.Post) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
}
