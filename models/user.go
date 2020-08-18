package models

import (
	"time"

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
}

// BeforeCreate hook runs before creating the user
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
