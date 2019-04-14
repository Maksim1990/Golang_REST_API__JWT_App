package post

import (
	"encoding/json"
	"fmt"
	"github.com/goRESTapi/database"
	"github.com/goRESTapi/jwtAuth"
	"github.com/goRESTapi/models"
	"github.com/goRESTapi/output"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetPostList(w http.ResponseWriter, req *http.Request) {
	db := database.DBConn()
	selDB, err := db.Query("SELECT id,user_id,title,description FROM posts ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	post := models.Post{}
	res := []models.Post{}
	for selDB.Next() {
		var id, user_id int
		var title, description string
		err = selDB.Scan(&id, &user_id, &title, &description)
		if err != nil {
			panic(err.Error())
		}
		post.ID = id
		post.UserId = user_id
		post.Title = title
		post.Description = description
		res = append(res, post)
	}

	//-- Generate JSON data list response
	output.JSONListResponse(w, res)

	defer db.Close()
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
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
		output.ExceptionMessage(w, "Post title can't be empty", 400)
	case err != nil:
		output.ExceptionMessage(w, err.Error(), 404)
	default:
		var userId int
		err = db.QueryRow("SELECT user_id FROM posts WHERE id=?", postId).Scan(&userId)
		if err != nil {
			panic(err.Error())
		} else {
			_, err = db.Query("UPDATE posts SET title=?,description=? WHERE id=?", title, description, postId)
			if err != nil {
				panic(err.Error())
			}
			id, err := strconv.Atoi(postId)
			if err != nil {
				panic(err.Error())
			}
			post.ID = id
			post.UserId = userId
			post.Title = title
			post.Description = description

			//-- Generate JSON response
			output.JSONResponse(w, post)
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
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				tokenData, _ := jwtAuth.ExtractClaims(bearerToken[1])
				userId := tokenData["id"]
				res, err := stmt.Exec(userId, title, description)
				if err != nil {
					panic(err.Error())
				}

				post := models.Post{}
				postId, _ := res.LastInsertId()
				post.ID = int(postId)
				post.Title = title
				post.UserId = int(userId.(float64))
				post.Description = description

				//-- Generate JSON response
				output.JSONResponse(w, post)
			}
		}

	}

	defer db.Close()
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	db := database.DBConn()
	selDB, err := db.Query("SELECT title FROM posts WHERE id=?", postId)
	if err != nil {
		panic(err.Error())
	} else {
		selPostRange := selDB.Next()
		if !selPostRange {
			output.ExceptionMessage(w, fmt.Sprintf("Post with ID %v was not found", postId), 404)
		} else {
			_, err := db.Query("DELETE FROM posts WHERE id=?", postId)
			if err != nil {
				panic(err.Error())
			}
			output.HttpResponse(w, fmt.Sprintf("Post with ID %v was successfully deleted", postId), 200)
		}
	}

	defer db.Close()
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]
	db := database.DBConn()
	selDB, err := db.Query("SELECT user_id,title,description FROM posts WHERE id=?", postId)
	if err != nil {
		panic(err.Error())
	}
	post := models.Post{}
	for selDB.Next() {
		var userId int
		var title, description string
		err = selDB.Scan(&userId, &title, &description)
		if err != nil {
			panic(err.Error())
		}
		post.Title = title
		post.Description = description
		post.UserId = userId

		id, err := strconv.Atoi(postId)
		if err != nil {
			panic(err.Error())
		}
		post.ID = id
	}

	//-- Generate JSON response
	output.JSONResponse(w, post)

	defer db.Close()

}
