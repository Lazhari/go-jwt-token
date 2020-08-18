package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/lazhari/web-jwt/models"

	"github.com/jinzhu/gorm"
	// Gorm postgres connector
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// ConnectDB establish the database connection
func ConnectDB() *gorm.DB {
	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)

	// Establish connection
	db, err := gorm.Open("postgres", dbConn)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
