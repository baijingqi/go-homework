package main

import (
    "fmt"
    "sync"
    "time"
)

type limitManager struct {
    maxTicketNum       int
    serverIngTicketNum int
    windowStartTime    int64
    lock               sync.RWMutex
}

func main() {
    limitManager := &limitManager{maxTicketNum: 20, windowStartTime: time.Now().Unix()}

    for i := 0; i < 60; i++ {
        ticket := limitManager.getTicket()
        if !ticket {
            fmt.Println("令牌数已超出")
        } else {
            fmt.Println("正常请求")
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
    if now-l.windowStartTime >= 60 {
        l.windowStartTime = now + 1
    }
    return true
}
