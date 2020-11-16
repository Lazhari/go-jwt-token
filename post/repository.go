package post

import "github.com/lazhari/web-jwt/models"

type Repository interface {
	CreatePost(*models.Post) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
}
