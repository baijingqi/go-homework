package main

import (
    "comment/internal/conf"
    "comment/internal/data"
    "context"
    "encoding/json"
    "flag"
    "fmt"
    "github.com/Shopify/sarama"
    "github.com/go-kratos/kratos/v2/config"
    "github.com/go-kratos/kratos/v2/config/file"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/tracing"
    "os"
    "strconv"
    "time"
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
    ctx := context.Background()

    defer c.Close()

    if err := c.Load(); err != nil {
        panic(err)
    }
    var bc conf.Bootstrap
    if err := c.Scan(&bc); err != nil {
        panic(err)
    }
    dataRepo, cleanup, err := data.NewData(bc.Data, logger)
    if err != nil {
        panic(err)
    }
    defer cleanup()

    //创建新的消费者
    consumer, err := sarama.NewConsumer([]string{bc.GetData().Kafka.GetAddr()}, nil)
    if err != nil {
        panic(err)
    }
    //根据topic获取所有的分区列表
    partitionList, err := consumer.Partitions("comment")
    if err != nil {
        fmt.Println("fail to get list of partition,err:", err)
    }

    //遍历所有的分区
    for p := range partitionList {
        //针对每一个分区创建一个对应分区的消费者
        pc, err := consumer.ConsumePartition("comment", int32(p), sarama.OffsetNewest)
        if err != nil {
            fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
            continue
        }
        fmt.Println(pc)

        for msg := range pc.Messages() {
            fmt.Printf("partition:%d Offse:%d Key:%v Value:%s \n",
                msg.Partition, msg.Offset, msg.Key, msg.Value)
            g := &data.Comment{}
            err := json.Unmarshal(msg.Value, &g)
            if err != nil {
                logger.Log(log.LevelError, err)
            }

            replyToUid := uint64(0)
            belongCommentId := uint64(0)
            parentInfo, err := dataRepo.DB.Comment.Get(ctx, g.ParentId)
            fmt.Println(parentInfo)
            if parentInfo != nil {
                replyToUid = parentInfo.UID
                belongCommentId = parentInfo.UID
            }
            fmt.Println("replyToUid=", replyToUid)
            save, err := dataRepo.DB.Comment.Create().
                SetUID(g.Uid).
                SetParentID(g.ParentId).
                SetRelationID(g.RelationId).
                SetRelationType(g.RelationType).
                SetBelongCommentID(belongCommentId).
                SetReplyToUID(replyToUid).
                SetContent(g.Content).
                SetCreatedAt(time.Now()).
                Save(ctx)

            if err != nil {
                logger.Log(log.LevelDebug, save, err)
                continue
            }

            if belongCommentId != 0 {
                dataRepo.Redis.HIncrBy("commentReplyNum", strconv.FormatUint(belongCommentId, 10), 1)
            }

            res, err := dataRepo.DB.CommentCount.Create().SetID(save.ID).SetCreatedAt(time.Now()).Save(ctx)
            if err != nil {
                logger.Log(log.LevelDebug, res, err)
                continue
            }
            logger.Log(log.LevelDebug, save, err)

            dataRepo.Redis.HIncrBy("commentCount",
                strconv.FormatUint(g.RelationId, 10)+"-"+strconv.FormatUint(uint64(g.RelationType), 10),
                1)
        }
    }
}
