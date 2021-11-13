package limit

import (
    "sync"
    "time"
)

type limitManager struct {
    maxTicketNum          uint           //单位时间内最大令牌数
    countSecond           uint           //单位时间 秒
    currentTicketNum      uint           //当前单位时间内请求令牌数
    perSecondGetTicketNum map[int64]uint //每一秒获取的令牌数
    seconds               []int64
    lastRequestTime       int64 //上次请求时间戳
    lock                  sync.RWMutex
}

func New(maxTicketNum uint, countSecond uint) *limitManager {
    return &limitManager{
        maxTicketNum:          maxTicketNum,
        countSecond:           countSecond,
        seconds:               make([]int64, countSecond),
        perSecondGetTicketNum: map[int64]uint{},
    }
}

func (l *limitManager) GetTicket() bool {
    now := time.Now().Unix()

    if now-l.lastRequestTime >= int64(l.countSecond) {
        l.perSecondGetTicketNum = make(map[int64]uint)
        l.seconds = make([]int64, l.countSecond)
        l.currentTicketNum = 0
    } else {
        if l.currentTicketNum+1 > l.maxTicketNum {
            return false
        }
    }

    l.lock.Lock()
    l.currentTicketNum += 1
    l.perSecondGetTicketNum[now] ++
    l.lastRequestTime = now

    //当前秒数据已经初始化
    if _, ok := l.perSecondGetTicketNum[now]; !ok {
        secondsLen := len(l.seconds)

        l.perSecondGetTicketNum[now] += 1
        //未到临界
        if uint(secondsLen) < l.countSecond {
            l.seconds = append(l.seconds, now)
        } else {
            //达到临界，需要将60秒以前的数据移除
            countStartTime := now - int64(l.countSecond)

            for _, value := range l.seconds {
                if value < countStartTime {
                    l.seconds = l.seconds[1:]
                    l.currentTicketNum -= l.perSecondGetTicketNum[value]
                    delete(l.perSecondGetTicketNum, value)
                }
            }
        }
    }

    l.lock.Unlock()
    return true
}
