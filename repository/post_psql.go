package repository

import "github.com/lazhari/web-jwt/models"

func (pr postgresRepository) CreatePost(p *models.Post) (*models.Post, error) {
	dbc := pr.db.Create(p)

	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return p, nil
}

func (pr postgresRepository) GetAllPosts() ([]models.Post, error) {
	return nil, nil
}
