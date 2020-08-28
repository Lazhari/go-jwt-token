package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Phone     string     `json:"phone"`
	IsActive  bool       `gorm:"DEFAULT:false" json:"isActive"`
	JobTitle  string     `json:"jobTitle"`
	Company   string     `json:"company"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `json:"password"`
	Posts     string     `gorm:"foreignKey:AuthorId" json:"posts"`
}

// BeforeCreate hook runs before creating the user
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

// Validate validation the user
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		// FirstName is require
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 40)),
		validation.Field(&u.LastName, validation.Required, validation.Length(3, 40)),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 20)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
