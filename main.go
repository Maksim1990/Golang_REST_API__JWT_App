package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goRESTapi/jwtAuth"
	"github.com/goRESTapi/models/post"
	"github.com/goRESTapi/models/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)


//-- Set middleware for requests
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	fmt.Println("Starting the application...")
	router.HandleFunc("/login", jwtAuth.GetAuthenticationToken).Methods("POST")
	router.HandleFunc("/register", user.RegisterUser).Methods("POST")

	router.HandleFunc("/users", jwtAuth.ValidateMiddleware(user.GetUserList)).Methods("GET")
	router.HandleFunc("/users/{id}", jwtAuth.ValidateMiddleware(user.GetUser)).Methods("GET")
	router.HandleFunc("/users/{id}", jwtAuth.ValidateMiddleware(user.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/users/{id}/update", jwtAuth.ValidateMiddleware(user.UpdateUser)).Methods("PUT")

	router.HandleFunc("/posts", jwtAuth.ValidateMiddleware(post.CreatePost)).Methods("POST")

	http.ListenAndServe(":9090", router)
}
