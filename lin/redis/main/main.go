package main

import (
	"CodeLin/lin/redis/distrbute_lock"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	//message_queue_list.ProducerMessageList()
	//message_queue_list.ConsumerMessageList()
	//go delay_queue.ProducerDelayMessage()
	//go delay_queue.ConsumerDelayMessage()
	//go message_queue_pubsub.Producer()
	//go message_queue_pubsub.Consumer1()
	//go message_queue_pubsub.Consumer2()
	lockTest()
}

func lockTest() {
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		nanoId := strconv.FormatInt(time.Now().UnixNano(), 10) // 模拟获取唯一标识
		for i := 0; i < 50; i++ {
			distrbute_lock.Lock(nanoId)
			count++
			distrbute_lock.UnLock(nanoId)
		}
		wg.Done()
	}()
	go func() {
		nanoId := strconv.FormatInt(time.Now().UnixNano(), 10) // 模拟获取唯一标识
		for i := 0; i < 50; i++ {
			distrbute_lock.Lock(nanoId)
			count++
			distrbute_lock.UnLock(nanoId)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(count)
}
