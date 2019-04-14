package jwtAuth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goRESTapi/database"
	"github.com/goRESTapi/models"
	"github.com/goRESTapi/output"
	"github.com/gorilla/context"
	"net/http"
	"os"
	"strings"
	"time"
)



//-- Duration of validation period of JWT Token
var expireTokenTime = 3600

//-- Get JWT secret phrase
var jwtSecret = os.Getenv("JWT_SECRET")

func GetAuthenticationToken(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user models.UserSecure
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	username := user.Username
	password := user.Password


	db := database.DBConn()

	err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	switch {
	case username == "":
		output.ExceptionMessage(w, "User name can't be empty ", 400)
	case password == "":
		output.ExceptionMessage(w, "Password can't be empty ", 400)
	case err == sql.ErrNoRows:
		output.ExceptionMessage(w, "User not found", 404)
	default:

		ttl := time.Duration(expireTokenTime) * time.Second
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"password": user.Password,
			"exp":      time.Now().UTC().Add(ttl).Unix(),
		})
		tokenString, error := token.SignedString([]byte(jwtSecret))
		if error != nil {
			fmt.Println(error)
		}
		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
	}
	defer db.Close()
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(jwtSecret), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(models.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(models.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(models.Exception{Message: "An authorization header is required"})
		}
	})
}
