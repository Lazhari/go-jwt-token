package driver

import (
	"log"

	"github.com/lazhari/web-jwt/models"

	"github.com/jinzhu/gorm"
	// Gorm postgres connector
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// ConnectDB establish the database connection
func ConnectDB() *gorm.DB {
	// pgURL, err := pq.ParseURL(os.Getenv("DB_URI"))

	// Establish connection
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=jwt-test password=toor sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
