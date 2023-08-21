package main

import (
	"fmt"
	"runtime"
)

// go 1.13
func main() {
	runtime.GOMAXPROCS(1)
	n := 0
	go func() {
		for {
			n++
			if n%10000000 == 0 {
				fmt.Println(n)
			}
		}
	}()
	for {
	}
}
