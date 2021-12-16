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
)

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

    repo := data.NewCommentRepo(dataData, logger)
    res, err := repo.CommentList(ctx, 1, 0, 0, 0, 1, 3)
    fmt.Println(res, err)

    return

}
