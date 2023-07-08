package lin

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// go redis 实现的轻量级消息队列2
type (
	PubSubMQ struct {
		rdb *redis.Client
	}

	Message struct {
		Topic     string
		Partition int
		Body      []byte
	}

	MessageHandler func(message *Message)
)

func NewPubSubMQ() *PubSubMQ {
	return &PubSubMQ{
		rdb: rdb,
	}
}

func (o *PubSubMQ) SendMsg(ctx context.Context, msg *Message) error {
	return o.rdb.Publish(ctx, o.partitionTopic(msg.Topic, msg.Partition), msg.Body).Err()
}

func (o *PubSubMQ) partitionTopic(topic string, partition int) string {
	return fmt.Sprintf("%s:%d", topic, partition)
}

func (o *PubSubMQ) ConsumeMsg(ctx context.Context, topic string, partition int, h MessageHandler) error {
	channel := o.rdb.Subscribe(ctx, o.partitionTopic(topic, partition)).Channel()
	for message := range channel {
		h(&Message{
			Topic:     topic,
			Partition: partition,
			Body:      []byte(message.Payload),
		})
	}
	return errors.New("channel closed")
}
