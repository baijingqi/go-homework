package dao

import (
    "database/sql"
    "fmt"
    "homework/database"
)

type User struct {
    id   int
    name string
    sex  int
}

func Find(uid int) *User {
    db := database.GetDB()
    user := User{}
    err := db.QueryRow("select * from `user` where id = ? ", uid).Scan(&user.id, &user.name, &user.sex)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil
        } else {
            fmt.Println("record Find function error")
        }
    }
    return &user
}

func GetList() []User {
    db := database.GetDB()
    rows, err := db.Query("select * from `user` where id >3 ")
    if err != nil {
        fmt.Println("query error:", err)
    }

    data := make([]User, 0)
    for rows.Next() {
        var user User
        err := rows.Scan(&user.id, &user.name, &user.sex)
        if err != nil {
            fmt.Println("data scan error :", err)
            continue
        }
        data = append(data, user)
    }
    return data
}
