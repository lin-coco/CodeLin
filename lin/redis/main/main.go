package main

import "CodeLin/lin/redis/message_queue_pubsub"

func main() {
	//message_queue_list.ProducerMessageList()
	//message_queue_list.ConsumerMessageList()
	//go delay_queue.ProducerDelayMessage()
	//go delay_queue.ConsumerDelayMessage()
	go message_queue_pubsub.Producer()
	go message_queue_pubsub.Consumer1()
	go message_queue_pubsub.Consumer2()
	select {}
}
