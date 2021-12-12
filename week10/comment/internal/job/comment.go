package main

import (
    "fmt"

    "github.com/Shopify/sarama"
)


func main() {
    config := sarama.NewConfig()
    //设置
    //ack应答机制
    config.Producer.RequiredAcks = sarama.WaitForAll

    //发送分区
    config.Producer.Partitioner = sarama.NewRandomPartitioner

    //回复确认
    config.Producer.Return.Successes = true

    //构造一个消息
    msg := &sarama.ProducerMessage{}
    msg.Topic = "comment"
    msg.Value = sarama.StringEncoder("test:comment device")

    //连接kafka
    client, err := sarama.NewSyncProducer([]string{"XTZJ-20211109XF.localdomain:9092"}, config)
    if err != nil {
        fmt.Println("producer closed,err:", err)
    }
    defer client.Close()

    //发送消息
    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send msg failed,err:", err)
        return
    }
    fmt.Printf("pid:%v offset:%v \n ", pid, offset)

}
