package delay_queue

import (
	redis2 "CodeLin/lin/redis"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

func ProducerDelayMessage() {
	for i := 0; i < 2; i++ {
		score := time.Now().Add(time.Second * 10).Unix()
		member := "member" + strconv.Itoa(i)
		redis2.Rdb.ZAdd(context.Background(), redis2.DelayKey, redis.Z{
			Score:  float64(score),
			Member: member,
		})
		log.Printf("插入延时消息:%s", member)
	}
}
