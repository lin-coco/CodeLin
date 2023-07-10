package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:2048")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("0.0.0.0:2048 开始监听...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			bytes := make([]byte, 1000)
			for {
				n, err := conn.Read(bytes)
				if err != nil {
					log.Println(err)
					break
				}
				fmt.Printf("server收到 %d 个字节，内容为：%s\n", n, bytes)
			}
		}(conn)
	}

}
