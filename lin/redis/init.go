package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Rdb            *redis.Client
	QueueListKey   = "queue:message"
	DelayKey       = "ZSet:message"
	QueuePubSubKey = "pubsub:message"
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "10.211.55.11:6379",
		Password: "syr1120@xys.com",
	})
	result, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("连接失败", err)
	}
	log.Println(result)

}
