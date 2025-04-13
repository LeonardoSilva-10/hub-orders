package db

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
    connStr := "host=localhost port=5432 user=postgres password=your-password dbname=hub_orders sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
}
