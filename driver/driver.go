package driver

import (
	"fmt"
	"os"

	"github.com/lazhari/web-jwt/models"

	"github.com/jinzhu/gorm"
	// Gorm postgres connector
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// ConnectDB establish the database connection
func ConnectDB() (*gorm.DB, error) {
	dbURI := os.Getenv("DB_URI")

	// Establish connection
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Post{})

	return db, nil
}
