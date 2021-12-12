package main

import (
    "comment/internal/conf"
    "comment/internal/data"
    "context"
    "flag"
    "fmt"
    "github.com/go-kratos/kratos/v2/config"
    "github.com/go-kratos/kratos/v2/config/file"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/tracing"
    "os"
    "strconv"
)

func main() {
    logger := log.With(log.NewStdLogger(os.Stdout),
        "ts", log.DefaultTimestamp,
        "caller", log.DefaultCaller,
        "trace_id", tracing.TraceID(),
        "span_id", tracing.SpanID(),
    )
    var flagconf string
    flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")

    c := config.New(
        config.WithSource(
            file.NewSource(flagconf),
        ),
    )

    defer c.Close()

    if err := c.Load(); err != nil {
        panic(err)
    }
    var bc conf.Bootstrap
    if err := c.Scan(&bc); err != nil {
        panic(err)
    }
    dataData, cleanup, err := data.NewData(bc.Data, logger)
    if err != nil {
        panic(err)
    }
    defer cleanup()
    cursor := uint64(0)
    ctx := context.Background()
    for{
        cmd := dataData.Redis.HScan("commentReplyNum", cursor, "*",1)
        page,cursor, _ :=cmd.Result()
        fmt.Println("page =", page)
        pageLen := len(page)
        for i:=0;i<pageLen;i+=2{
            id, err := strconv.ParseUint(page[i], 10, 64)
            if err != nil {
                continue
            }
            replyNum, err := strconv.ParseUint(page[i], 10, 64)
            if err != nil {
                continue
            }
            dataData.DB.CommentCount.UpdateOneID(id).AddReplyNum(uint32(replyNum)).Exec(ctx)
        }
        if cursor == 0 {
            break
        }
    }


}
