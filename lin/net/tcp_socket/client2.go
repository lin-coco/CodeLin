package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:2048")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		n, _ := conn.Write([]byte("hellloo"))
		fmt.Printf("client2 发送了 %d 字节数据\n", n)
		time.Sleep(time.Second * 2)
	}
}
