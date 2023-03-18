package main

import (
    "database/sql"
    "fmt"
    "log"
    
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config {
        User:   "root",
        Passwd: "toor",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "groceries",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}