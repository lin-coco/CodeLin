package message_queue_pubsub

import (
	"CodeLin/lin/redis"
	"context"
	"log"
	"strconv"
	"time"
)

func Consumer1() {
	subscribe := redis.Rdb.Subscribe(context.Background(), redis.QueuePubSubKey)
	channel := subscribe.Channel()
	for message := range channel {
		a, _ := strconv.Atoi(message.Payload)
		log.Printf("Consumer1 订阅到了消息:%d", a)
		time.Sleep(2 * time.Second)
	}
	log.Println("Consumer1 退出")
}

func Consumer2() {
	subscribe := redis.Rdb.Subscribe(context.Background(), redis.QueuePubSubKey)
	channel := subscribe.Channel()
	for message := range channel {
		a, _ := strconv.Atoi(message.Payload)
		log.Printf("Consumer2 订阅到了消息:%d", a)
		time.Sleep(2 * time.Second)
	}
	log.Println("Consumer2 退出")
}
