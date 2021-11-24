package main

import (
    "fmt"
     "github.com/pkg/errors"
)

func main()  {
    err := errors.New("xxx")
    fmt.Println(err)
    w := errors.Wrap(err,"Wrap了一个错误")
    fmt.Println(w)
}