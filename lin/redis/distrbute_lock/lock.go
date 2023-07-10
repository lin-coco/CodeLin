package distrbute_lock

import (
	"CodeLin/lin/redis"
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"log"
	"time"
)

var (
	channel = make(chan struct{}, 0)
)

func Lock(id string) {
	for {
		success, err := redis.Rdb.SetNX(context.Background(), redis.DistrbutedLockKey, id, time.Second*10).Result()
		if err != nil {
			// 抢锁失败，继续自旋
			log.Println(id, "抢锁失败，继续自旋")
			continue
		}
		if success {
			// 抢锁成功
			log.Println(id, "抢锁成功")
			go watchDog(id)
			break
		}
	}
}

// UnLock 解锁过程原子化，防止在两条redis语句之前误删
func UnLock(id string) {
	script := redis2.NewScript(`
		if redis.call('get',KEYS[1]) == ARGV[1]
		then 
			return redis.call('del',KEYS[1])
		else
			return 0
		end
`)
	result, err := script.Run(context.Background(), redis.Rdb, []string{redis.DistrbutedLockKey}, id).Result()
	if err != nil || result == 0 {
		fmt.Println("unlock failed", err)
	}
	log.Println(id, "解锁成功")
	channel <- struct{}{}
}

// watchDog 看门狗实现，实现锁的续期。在加锁后启动watchDog，解锁后退出watchDog，watchDow需要定期的延长key的过期时间
func watchDog(id string) {
	ticker := time.NewTicker(8 * time.Second)
	script := redis2.NewScript(`
		if redis.call('get',KEYS[1]) == ARGV[1]
		then
			return redis.call('expire',KEYS[1],ARGV[2])
		else
			return 0
		end
`)
	for {
		select {
		case <-ticker.C:
			result, err := script.Run(context.Background(), redis.Rdb, []string{redis.DistrbutedLockKey}, id, 8).Result()
			if err != nil || result == 0 {
				log.Println("expire failed", err)
			}
		case <-channel:
			log.Println("已经解锁，看门狗退出")
			return
		}
	}
}
