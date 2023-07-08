package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:2048")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello"))
	fmt.Printf("client 发送了 %d 字节数据", n)
}
