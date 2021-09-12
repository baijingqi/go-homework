package main

import (
    "fmt"
    "homework/dao"
    "homework/database"
)

func main() {
    database.InitDB()
    res := dao.Find(3)
    fmt.Println("单条查询", res)

    data := dao.GetList()
    fmt.Println(data)
}
