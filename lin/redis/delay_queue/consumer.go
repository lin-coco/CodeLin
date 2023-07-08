package delay_queue

import (
	"CodeLin/lin/redis"
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConsumerDelayMessage() {
	for {
		values, err := redis.Rdb.ZRangeByScore(context.Background(), redis.DelayKey, &redis2.ZRangeBy{
			Min:    "0",
			Max:    fmt.Sprint(time.Now().Unix()),
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		if len(values) == 0 {
			log.Println("没有延时消息")
			time.Sleep(time.Second)
			continue
		}
		value := values[0]
		result, err := redis.Rdb.ZRem(context.Background(), redis.DelayKey, value).Result()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		if result == 1 {
			log.Println("消费到数据：", value, "当前时间是：", time.Now().Unix())
			// 模拟一个耗时的操作
			time.Sleep(2 * time.Second)
		}
	}
}
