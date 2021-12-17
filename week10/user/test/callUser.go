package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    pb "user/api/user"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()

    c := pb.NewUserClient(conn)
    userIds := []uint64{1,2,3,5,6}
    req := &pb.UserListRequest{
        UserIds: userIds,
    }
    r, err := c.UserList(context.Background(), req)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(r.List)
    fmt.Println(r.Code)
    fmt.Println(r.Msg)
}

