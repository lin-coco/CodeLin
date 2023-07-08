package message_queue_list

import (
	"CodeLin/lin/redis"
	"context"
	"log"
	"math/rand"
	"time"
)

func ProducerMessageList() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		t := rand.Intn(1000)
		log.Printf("生产一条消息 message:%d", t)
		_, err := redis.Rdb.LPush(context.Background(), redis.QueueListKey, t).Result()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
