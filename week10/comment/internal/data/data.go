package data

import (
    "comment/ent"
    "comment/internal/conf"
    _ "database/sql"
    "fmt"
    "github.com/Shopify/sarama"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis"
    _ "github.com/go-sql-driver/mysql"
    "github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCommentRepo)

// Data .
type Data struct {
    db    *ent.Client
    Bb    *ent.Client
    Kafka sarama.SyncProducer
    log   *log.Helper
    DB    *ent.Client
    Redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    client, err := ent.Open(c.Database.GetDriver(), c.Database.GetSource())
    if err != nil {
        fmt.Println("初始化 数据库 失败", err)
        return nil, nil, err
    }

    kafkaClient, err := initKafka(c, logger)
    if err != nil {
        fmt.Println("初始化kafka 失败", err)
        logger.Log(log.LevelFatal, err)
        return nil, nil, err
    }

    redisClient, err := initRedis(c, logger)
    if err != nil {
        fmt.Println("初始化 redis 失败", err)
        logger.Log(log.LevelFatal, err)
        return nil, nil, err
    }

    return &Data{
        db:    client,
        DB:    client,
        Kafka: kafkaClient,
        Redis: redisClient,
    }, cleanup, nil
}

func initKafka(c *conf.Data, logger log.Logger) (sarama.SyncProducer, error) {
    config := sarama.NewConfig()
    //设置
    //ack应答机制
    config.Producer.RequiredAcks = sarama.WaitForAll
    //发送分区
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    //回复确认
    config.Producer.Return.Successes = true
    //连接kafka
    client, err := sarama.NewSyncProducer([]string{c.Kafka.GetAddr()}, config)
    if err != nil {
        return nil, err
    }

    return client, nil
}

func initRedis(c *conf.Data, logger log.Logger) (*redis.Client, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     c.Redis.GetAddr(),
        DB:       0,
    })

    _, err := client.Ping().Result()
    return client, err

}
