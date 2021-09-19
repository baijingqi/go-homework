package main

import (
    "context"
    "fmt"
    "github.com/neilotoole/errgroup"
    "net/http"
    "os"
    "os/signal"
    "time"
)

type myHttp struct {
    ctx context.Context
    g   *errgroup.Group
}

func user(w http.ResponseWriter, r *http.Request, g *errgroup.Group) {
    fmt.Fprintf(w, "you called user func")
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("模拟后台耗时任务")
            time.Sleep(3 * time.Second)
        }
    }()
}

func hello(w http.ResponseWriter, r *http.Request, g *errgroup.Group) {
    fmt.Fprintf(w, "you called hello func")

}
func (myHttp *myHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("接收到请求", r.URL.Path)

    switch r.URL.Path {
    case "/user":
        user(w, r, myHttp.g)
    case "/hello":
        hello(w, r, myHttp.g)
    default:
        fmt.Fprintf(w, "接收到请求"+r.URL.Path)
    }
}

func main() {
    g, ctx := errgroup.WithContext(context.Background())

    g.Go(func() error {
        fmt.Println(456)

        sigs := make(chan os.Signal)
        signal.Notify(sigs)
        // 打印进程id
        fmt.Println("PID:", os.Getpid())
        select {
        case s := <-sigs:
            fmt.Println("usr2 signal", s)
            fmt.Println("收到自定义退出信号2")
            return fmt.Errorf("自定义退出2")
        }
    })

    g.Go(func() error {
        fmt.Println(123)
        return http.ListenAndServe(":9999", &myHttp{ctx: ctx, g: g})
    })
    fmt.Println("123123")
    err := g.Wait()
    if err != nil {
        return
    }
}
