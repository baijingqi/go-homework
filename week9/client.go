package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "io/ioutil"
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
func main()  {

    //writeFile()
    //return
    data, err := ioutil.ReadFile("./a")
    if err!=nil{
        panic("读取文件失败")
    }


    //fmt.Println(BytesToInt32(data[0:4]))
    //fmt.Println((data[4:6]))
    //fmt.Println(BytesToInt16(data[4:6]))
    //
    //return
    conn, err := net.Dial("tcp", "0.0.0.0:10086")
    if err!= nil{
        panic("连接失败")
    }
    for i:=0;i<5;i++{
        _, e := conn.Write(data)
        if e != nil {
            fmt.Println("Error to send message because of ", e.Error())
        }
    }
}

func writeFile() {
    packageData := goim{
       PacketLen: 53,                                  //4
       HeaderLen: 16,                                  //2
       Version:   1,                                   //2
       Operation: 1,                                   //4
       Sequence:  1,                                   //4
       Body:      []byte("token:123123123;uid:11019;name:大白"), //25
    }
    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, packageData.PacketLen)
    binary.Write(bytesBuffer, binary.BigEndian, packageData.HeaderLen)
    binary.Write(bytesBuffer, binary.BigEndian, packageData.Version)
    binary.Write(bytesBuffer, binary.BigEndian, packageData.Operation)
    binary.Write(bytesBuffer, binary.BigEndian, packageData.Sequence)
    binary.Write(bytesBuffer, binary.BigEndian, packageData.Body)
    //fmt.Println("=====%b,输出二进制====")

    //fmt.Printf("%b\n", bytesBuffer)
    fmt.Printf("%d\n", bytesBuffer.Bytes())
    ioutil.WriteFile("b", bytesBuffer.Bytes(), 0777)
}