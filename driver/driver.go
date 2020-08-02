package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

// ConnectDB establish the database connection
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("DB_URI"))

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
