package database

import (
    "os"
    "database/sql"
    _ "github.com/lib/pq"
)

func DBConn() (db *sql.DB) {
    dbDriver := os.Getenv("DB_PG_ENGINE")
    dbUser := os.Getenv("DB_PG_USER")
    dbPass := os.Getenv("DB_PG_PASSWORD")
    dbName := os.Getenv("DB_PG_NAME")
    connStr := dbDriver+"://"+dbUser+":"+dbPass+"@db:5432/"+dbName+"?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    return db
}
