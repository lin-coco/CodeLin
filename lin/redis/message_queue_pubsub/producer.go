package message_queue_pubsub

import (
	"CodeLin/lin/redis"
	"context"
	"log"
	"math/rand"
	"time"
)

func Producer() {
	time.Sleep(3 * time.Second)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		n := rand.Intn(1000)
		_, err := redis.Rdb.Publish(context.Background(), redis.QueuePubSubKey, n).Result()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("发布了消息:%d", n)
		time.Sleep(time.Second)
	}
}
