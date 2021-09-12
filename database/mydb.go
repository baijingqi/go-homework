package database

import (
    "database/sql"
    _ "database/sql"
    _ "time"

    _ "github.com/go-sql-driver/mysql"
)

var instance *sql.DB

func InitDB() {
    dsn := "bai:bai@tcp(192.168.10.10:3306)/homestead?charset=utf8"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    instance = db
    return
}

func GetDB() *sql.DB {
    return instance
}
