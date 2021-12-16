package main

import (
    pb "comment/api/comment"
    "context"
    "fmt"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()

    c := pb.NewCommentClient(conn)

    req := &pb.CommentListRequest{
        RelationId:1,
        RelationType:1,
        Page:1,
    }
    r, err := c.CommentList(context.Background(), req)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(r.List)
    fmt.Println(r.Code)
    fmt.Println(r.Msg)
}

