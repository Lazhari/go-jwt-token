package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lazhari/web-jwt/controllers"
	"github.com/lazhari/web-jwt/driver"
	"github.com/lazhari/web-jwt/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var db *gorm.DB

// Route represent the application routes
type Route struct {
	path    string
	handler http.HandlerFunc
	method  string
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)

	routes := []Route{
		{
			path:    "/sign-up",
			handler: controller.SignUpHandler(db),
			method:  "POST",
		},
		{
			path:    "/login",
			handler: controller.LoginHandler(db),
			method:  "POST",
		},
		{
			path:    "/protected",
			handler: middleware.IsAuthenticated(controller.ProtectedHandler()),
			method:  "GET",
		},
	}

	for _, route := range routes {
		r.HandleFunc(route.path, route.handler).Methods(route.method)
	}

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8050", r))
}
