package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lazhari/web-jwt/controllers"

	"github.com/lazhari/web-jwt/middleware"

	"github.com/lazhari/web-jwt/driver"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

func init() {
	gotenv.Load()
}

var db *sql.DB

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)
	r.HandleFunc("/sign-up", controller.SignUpHandler(db)).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler(db)).Methods("POST")
	r.HandleFunc("/protected", middleware.IsAuthenticated(controller.ProtectedHandler())).Methods("GET")

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8050", r))
}
