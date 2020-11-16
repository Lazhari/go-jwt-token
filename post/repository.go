package post

import "github.com/lazhari/web-jwt/models"

type Repository interface {
	Create(*models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
}
