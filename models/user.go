package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
