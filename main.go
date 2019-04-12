package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "github.com/joho/godotenv"
    "strconv"
    "log"
    "github.com/goRESTapi/database"
    "github.com/goRESTapi/output"
    "github.com/goRESTapi/jwtAuth"
)

type User struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func GetUserList(w http.ResponseWriter, req *http.Request) {
    db:=database.DBConn()
    selDB, err := db.Query("SELECT id,username,password FROM users ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    usr := User{}
    res := []User{}
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

    m := make(map[string][]User)
    m["data"] = res

    json.NewEncoder(w).Encode(m)

    defer db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  userId := params["id"]
  decoder := json.NewDecoder(r.Body)
  var user User
  err := decoder.Decode(&user)
  if err != nil {
      panic(err)
  }
  username := user.Username
  password := user.Password
  db:=database.DBConn()

  switch{
  case username == "":
    output.ExceptionMessage(w,"User name can't be empty",400)
  case password == "":
    output.ExceptionMessage(w,"Password can't be empty",400)
  case err !=nil:
    output.ExceptionMessage(w,err.Error(),404)
  default:
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err.Error())
    }
    selDB, err := db.Query("SELECT username,password FROM users WHERE id=?",userId)
    if err != nil {
        panic(err.Error())
    }else{

      selUserRange:=selDB.Next()
      if !selUserRange{
        output.ExceptionMessage(w,fmt.Sprintf("User with ID %v was not found",userId), 404)
      }else{
        _, err = db.Query("UPDATE users SET username=?, password=? WHERE id=?",username,hashedPassword,userId)
        if err != nil {
            panic(err.Error())
        }
        user := User{}
        id, err := strconv.Atoi(userId)
        if err != nil {
            panic(err.Error())
        }
        user.Id = id
        user.Username = username
        user.Password = string(hashedPassword)

        m := make(map[string]User)
        m["data"] = user
        json.NewEncoder(w).Encode(m)
      }
    }
  }

  defer db.Close()
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var user User
    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }
    username := user.Username
    password := user.Password
    db:=database.DBConn()

    err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
    switch {
    case username == "":
      output.ExceptionMessage(w,"User name can't be empty", 400)
    case password == "":
      output.ExceptionMessage(w,"Password can't be empty", 400)
    case err == sql.ErrNoRows:
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            panic(err.Error())
        }

        stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        res, err :=stmt.Exec(username, hashedPassword)
        if err != nil {
            panic(err.Error())
        }

        user := User{}
        userId, _ := res.LastInsertId()
        user.Id = int(userId)
        user.Username = username
        user.Password = string(hashedPassword)


        m := make(map[string]User)
        m["data"] = user
        json.NewEncoder(w).Encode(m)

    case err != nil:
      output.ExceptionMessage(w,err.Error(), 400)
    default:
      output.ExceptionMessage(w,"Bad request", 400)
    }

    defer db.Close()
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId := params["id"]
    db:=database.DBConn()
    selDB, err := db.Query("SELECT username,password FROM users WHERE id=?",userId)
    if err != nil {
        panic(err.Error())
    }else {
        selUserRange := selDB.Next()
        if !selUserRange {
            output.ExceptionMessage(w, fmt.Sprintf("User with ID %v was not found", userId), 404)
        } else {
            _, err := db.Query("DELETE FROM users WHERE id=?",userId)
            if err != nil {
                panic(err.Error())
            }
            output.HttpResponse(w,fmt.Sprintf("User with ID %v was successfully deleted",userId),200)
        }
    }

    defer db.Close()
}
func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userId := params["id"]
    db:=database.DBConn()
    selDB, err := db.Query("SELECT username,password FROM users WHERE id=?",userId)
    if err != nil {
        panic(err.Error())
    }
    user := User{}
    for selDB.Next() {
        var username,password string
        err = selDB.Scan(&password,&username)
        if err != nil {
           panic(err.Error())
        }
        user.Username = username
        user.Password = password

        id, err := strconv.Atoi(userId)
        if err != nil {
            panic(err.Error())
        }
        user.Id = id
    }

    m := make(map[string]User)
    m["data"] = user

    json.NewEncoder(w).Encode(m)

    defer db.Close()

}
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
    router.HandleFunc("/register", RegisterUser).Methods("POST")

    router.HandleFunc("/users", jwtAuth.ValidateMiddleware(GetUserList)).Methods("GET")
    router.HandleFunc("/users/{id}", jwtAuth.ValidateMiddleware(GetUser)).Methods("GET")
    router.HandleFunc("/users/{id}", jwtAuth.ValidateMiddleware(DeleteUser)).Methods("DELETE")
    router.HandleFunc("/users/{id}/update", jwtAuth.ValidateMiddleware(UpdateUser)).Methods("PUT")

    http.ListenAndServe(":9090", router)

}
