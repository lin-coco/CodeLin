package lin

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// go redis 实现的轻量消息队列
// List实现
type (
	MQ struct {
		Topic string
	}

	Msg struct {
		Body []byte
	}

	Handler func(msg *Msg) error

	Result struct {
		n   int
		err error
	}
)

var (
	rdb *redis.Client
)

func NewListMQ(key string) *MQ {
	return &MQ{Topic: key}
}

func (o *MQ) Produce(ctx context.Context, msg Msg) error {
	return rdb.LPush(ctx, o.Topic, msg.Body).Err()
}

func (o *MQ) Consume(ctx context.Context, h Handler) error {
	for {
		bytes, err := rdb.LIndex(ctx, o.Topic, -1).Bytes()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
		if errors.Is(err, redis.Nil) {
			time.Sleep(time.Second * 5)
			continue
		}
		// 处理消息
		err = h(&Msg{Body: bytes})
		if err != nil {
			continue
		}
		// 处理成功删除数据
		if err = rdb.RPop(ctx, o.Topic).Err(); err != nil {
			return err
		}
	}
}
