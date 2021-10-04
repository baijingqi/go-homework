package main

import (
    "fmt"
    "sync"
    "time"
)

type limitManager struct {
    maxTicketNum       int
    serverIngTicketNum int
    countSecond        int
    windowStartTime    int64
    lock               sync.RWMutex
}

func main() {
    limitManager := &limitManager{maxTicketNum: 20, countSecond: 20, windowStartTime: time.Now().Unix()}

    for i := 0; i < 100; i++ {
        ticket := limitManager.getTicket()
        if !ticket {
            fmt.Println("令牌数已超出")
        } else {
            fmt.Println("正常请求")
            go func() {
                defer limitManager.reduceServingTicket()
                fmt.Println("do sth")
            }()
        }
    }

}

func (l *limitManager) getTicket() bool {
    if l.serverIngTicketNum >= l.maxTicketNum {
        return false
    }

    l.lock.Lock()
    l.serverIngTicketNum += 1
    l.lock.Unlock()

    now := time.Now().Unix()
    if now-l.windowStartTime >= int64(l.countSecond) {
        l.windowStartTime = now + 1
    }
    return true
}

func (l *limitManager) reduceServingTicket() {
    l.serverIngTicketNum--
}
