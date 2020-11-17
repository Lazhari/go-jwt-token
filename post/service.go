package post

import "github.com/lazhari/web-jwt/models"

// Service the post service
type Service interface {
	Create(*models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
}
