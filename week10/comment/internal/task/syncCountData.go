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
    "sync"
)
var wg sync.WaitGroup

func main() {
    logger := log.With(log.NewStdLogger(os.Stdout),
        "ts", log.DefaultTimestamp,
        "caller", log.DefaultCaller,
        "trace_id", tracing.TraceID(),
        "span_id", tracing.SpanID(),
    )
    var flagconf string
    flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf config.yaml")

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
    ctx := context.Background()

    defer cleanup()
    scanFunc := func(key string,wg *sync.WaitGroup) {
        cursor := uint64(0)
        for {
            cmd := dataData.Redis.HScan(key, cursor, "*", 1)
            result, cursor, err := cmd.Result()
            fmt.Println("result = ", result)
            fmt.Println("cursor", cursor)
            fmt.Println("err  =", err)
            pageLen := len(result)
            for i := 0; i < pageLen; i += 2 {
                id, err := strconv.ParseUint(result[i], 10, 64)
                if err != nil {
                    continue
                }
                num, err := strconv.ParseUint(result[i+1], 10, 64)
                if err != nil {
                    continue
                }
                switch key {
                case "commentReplyNum":
                    dataData.DB.CommentCount.UpdateOneID(id).AddReplyNum(uint32(num)).Exec(ctx)
                case "commentDislikeNum":
                    dataData.DB.CommentCount.UpdateOneID(id).AddDislikeNum(uint32(num)).Exec(ctx)
                case "commentPraiseNum":
                    dataData.DB.CommentCount.UpdateOneID(id).AddPraiseNum(uint32(num)).Exec(ctx)
                }
                fmt.Println("id=",id," num=", num)
            }
            if cursor == 0 {
                wg.Done()
                break
            }
        }
    }
    wg.Add(3)
    go scanFunc("commentPraiseNum",&wg)
    go scanFunc("commentReplyNum", &wg)
    go scanFunc("commentDislikeNum",&wg)
    wg.Wait()
}
