package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post model
type Post struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	AuthorID  string     `json:"AuthorID"`
}

// BeforeCreate hooks run before create a post
func (p *Post) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

// Validate validation the post
func (p Post) Validate() error {
	return validation.ValidateStruct(&p,
		// FirstName is require
		validation.Field(&p.Title, validation.Required, validation.Length(3, 40)),
		validation.Field(&p.Body, validation.Required),
	)
}
