package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lazhari/web-jwt/api"
	"github.com/lazhari/web-jwt/middleware"
	"github.com/lazhari/web-jwt/post"
	"github.com/lazhari/web-jwt/repository"
	"github.com/lazhari/web-jwt/user"
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
	repo, err := repository.NewPostgreRepository()
	if err != nil {
		log.Fatal(err)
	}

	userService := user.NewAuthService(repo)
	postService := post.NewPostService(repo)
	userHandler := api.NewUserHandler(userService)
	postHandler := api.NewPostHandler(postService)

	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)

	routes := []Route{
		{
			path:    "/sign-up",
			handler: userHandler.SignUp,
			method:  "POST",
		},
		{
			path:    "/login",
			handler: userHandler.Login,
			method:  "POST",
		},
		{
			path:    "/posts",
			handler: middleware.IsAuthenticated(postHandler.CreatePost),
			method:  "POST",
		},
		{
			path:    "/posts",
			handler: middleware.IsAuthenticated(postHandler.GetAllPosts),
			method:  "GET",
		},
		// {
		// 	path:    "/posts/{id}",
		// 	handler: middleware.IsAuthenticated(controller.GetPostByID(db)),
		// 	method:  "GET",
		// },
	}

	for _, route := range routes {
		r.HandleFunc(route.path, route.handler).Methods(route.method)
	}

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8050"
	}

	errs := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on port http://localhost:%s\n", port)
		errs <- http.ListenAndServe(":"+port, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
