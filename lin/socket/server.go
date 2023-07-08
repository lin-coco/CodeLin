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
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			bytes := make([]byte, 5)
			n, err := conn.Read(bytes)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("server收到 %d 个字节，内容为：%s", n, bytes)
		}(conn)
	}

}
