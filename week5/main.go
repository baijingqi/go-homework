package main

import (
    "fmt"
    "time"
    mylimit "week5/limit"
)

func main() {
    //arr := []int{2,3,4,5,6}
    //
    //for key,value := range arr{
    //    fmt.Println("key=",key, "value=",value)
    //}
    //fmt.Println(arr[1:])

    limit := mylimit.New(5, 5)
    for {
        for i := 0; i < 10; i++ {
            fmt.Println(limit.GetTicket())
        }
        fmt.Println("")
        time.Sleep(4 * time.Second)
    }
}
