package message_queue_list

import (
	redis2 "CodeLin/lin/redis"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConsumerMessageList() {
	for {
		result, err := redis2.Rdb.BRPop(context.Background(), 5*time.Second, redis2.QueueListKey).Result()
		//result, err := rdb.RPop(context.Background(), queueListKey).Result()
		if err == redis.Nil {
			log.Println("未查询到消息")
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("消费消息：%s", result)
	}
}
