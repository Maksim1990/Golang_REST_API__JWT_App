package post

import (
	"encoding/json"
	"fmt"
	"github.com/goRESTapi/models"
	"github.com/goRESTapi/output"
	"github.com/goRESTapi/database"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func GetUserList(w http.ResponseWriter, req *http.Request) {
	db := database.DBConn()
	selDB, err := db.Query("SELECT id,username,password FROM users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	usr := models.User{}
	res := []models.User{}
	for selDB.Next() {
		var id int
		var username, password string
		err = selDB.Scan(&id, &username, &password)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Username = username
		usr.Password = password
		res = append(res, usr)
	}

	//-- Generate JSON data list response
	output.JSONListResponse(w,res)

	defer db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	username := user.Username
	password := user.Password
	db := database.DBConn()

	switch {
	case username == "":
		output.ExceptionMessage(w, "User name can't be empty", 400)
	case password == "":
		output.ExceptionMessage(w, "Password can't be empty", 400)
	case err != nil:
		output.ExceptionMessage(w, err.Error(), 404)
	default:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(err.Error())
		}
		selDB, err := db.Query("SELECT username,password FROM users WHERE id=?", userId)
		if err != nil {
			panic(err.Error())
		} else {

			selUserRange := selDB.Next()
			if !selUserRange {
				output.ExceptionMessage(w, fmt.Sprintf("User with ID %v was not found", userId), 404)
			} else {
				_, err = db.Query("UPDATE users SET username=?, password=? WHERE id=?", username, hashedPassword, userId)
				if err != nil {
					panic(err.Error())
				}
				user := models.User{}
				id, err := strconv.Atoi(userId)
				if err != nil {
					panic(err.Error())
				}
				user.Id = id
				user.Username = username
				user.Password = string(hashedPassword)

				//-- Generate JSON response
				output.JSONResponse(w, user)
			}
		}
	}

	defer db.Close()
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var post models.Post
	err := decoder.Decode(&post)
	if err != nil {
		panic(err)
	}
	title := post.Title
	description := post.Description
	db := database.DBConn()


	switch {
	case title == "":
		output.ExceptionMessage(w, "Title can't be empty", 400)
	default:
		stmt, err := db.Prepare("INSERT INTO posts(user_id,title, description) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		res, err := stmt.Exec(1,title, description)
		if err != nil {
			panic(err.Error())
		}

		post := models.Post{}
		postId, _ := res.LastInsertId()
		post.ID = int(postId)
		post.Title = title
		post.Description = description

		//-- Generate JSON response
		output.JSONResponse(w, post)
	}

	defer db.Close()
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	db := database.DBConn()
	selDB, err := db.Query("SELECT username,password FROM users WHERE id=?", userId)
	if err != nil {
		panic(err.Error())
	} else {
		selUserRange := selDB.Next()
		if !selUserRange {
			output.ExceptionMessage(w, fmt.Sprintf("User with ID %v was not found", userId), 404)
		} else {
			_, err := db.Query("DELETE FROM users WHERE id=?", userId)
			if err != nil {
				panic(err.Error())
			}
			output.HttpResponse(w, fmt.Sprintf("User with ID %v was successfully deleted", userId), 200)
		}
	}

	defer db.Close()
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	db := database.DBConn()
	selDB, err := db.Query("SELECT username FROM users WHERE id=?", userId)
	if err != nil {
		panic(err.Error())
	}
	user := models.User{}
	for selDB.Next() {
		var username string
		err = selDB.Scan(&username)
		if err != nil {
			panic(err.Error())
		}
		user.Username = username

		id, err := strconv.Atoi(userId)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
	}

	//-- Generate JSON response
	output.JSONResponse(w, user)

	defer db.Close()

}
