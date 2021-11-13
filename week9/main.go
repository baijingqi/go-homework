package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "net"
)

type goim struct {
    PacketLen int32
    HeaderLen int16
    Version   int16
    Operation int32
    Sequence  int32
    Body      []byte
}

func main() {
    address := net.TCPAddr{
        IP:   net.IP("0.0.0.0").To4(),
        Port: 10086,
    }
    listener, err := net.ListenTCP("tcp4", &address) // 创建TCP4服务器端监听器
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listener.AcceptTCP()
        if err != nil {
            log.Fatal(err) // 错误直接退出
        }
        fmt.Println("remote address:", conn.RemoteAddr())
        fmt.Println(conn)

        go handleRequest(conn)

    }
}

func handleRequest(conn *net.TCPConn)  {
    defer func(conn *net.TCPConn) {
        err := conn.Close()
        if err != nil {
            log.Fatal("关闭连接失败",err)
        }
    }(conn)

    buf := make([]byte, 16)
    bodyStr:= ""
    for {
        if  readLen, _ :=conn.Read(buf);readLen != 0{
            fmt.Println(buf)
            packageLen := BytesToInt32(buf[0:4])
            headerLen := BytesToInt16(buf[4:6])
            version := BytesToInt16(buf[6:8])
            operation := BytesToInt16(buf[8:12])
            sequence := BytesToInt32(buf[12:16])
            bodyLen := packageLen-int32(headerLen)
            fmt.Println("packageLen=",packageLen)
            fmt.Println("headerLen=",headerLen)
            fmt.Println("version=",version)
            fmt.Println("operation=",operation)
            fmt.Println("sequence=",sequence)
            fmt.Println("bodyLen=",bodyLen)

            var bodyCount int32
            bodyBuf := make([]byte, 12)
            for bodyCount < bodyLen{
                bodyReadLen, _ := conn.Read(bodyBuf)
                fmt.Println(bodyBuf)
                if bodyReadLen!=0 {
                    fmt.Println("bodyCount=",bodyCount, "bodyLen=", bodyLen)

                    if bodyLen-bodyCount > 12 {
                        fmt.Println(">1024 0",string(bodyBuf))
                        bodyStr += string(bodyBuf)
                        bodyCount += int32(bodyReadLen)
                    }else{
                        fmt.Println(">1024 0",string(bodyBuf))
                        fmt.Println("<1024 2 ",string(bodyBuf))
                        fmt.Println("<1024 3 ",string(bodyBuf[0:(bodyLen-bodyCount)]))
                        bodyStr += string(bodyBuf[0:(bodyLen-bodyCount)])
                        bodyCount += bodyLen - bodyCount
                    }
                }
                fmt.Println("bodyStr=",bodyStr)
            }
            break
        }
    }
}

func BytesToInt16(b []byte) int16 {
    bytesBuffer := bytes.NewBuffer(b)
    var x int16
    binary.Read(bytesBuffer, binary.BigEndian, &x)
    return x
}
func BytesToInt32(b []byte) int32 {
    bytesBuffer := bytes.NewBuffer(b)
    var x int32
    binary.Read(bytesBuffer, binary.BigEndian, &x)
    return x
}


