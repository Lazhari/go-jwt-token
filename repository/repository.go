package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lazhari/web-jwt/driver"
	"github.com/lazhari/web-jwt/post"
	"github.com/lazhari/web-jwt/user"
)

// Repository interface
type Repository interface {
	user.Repository
	post.Repository
}

type postgresRepository struct {
	db *gorm.DB
}

// NewPostgreRepository Create the postgres repository
func NewPostgreRepository() (Repository, error) {
	db, err := driver.ConnectDB()
	if err != nil {
		return nil, err
	}

	repo := postgresRepository{
		db,
	}
	return repo, nil
}
