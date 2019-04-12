package database

import (
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
    dbDriver := os.Getenv("DB_ENGINE")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(db:3306)/"+dbName)

    //-- For non docker environment
    //db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
