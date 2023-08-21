package main

import (
	"fmt"
	"runtime"
	"time"
)

// go 1.13 sigurg
func main() {
	var x int
	threads := runtime.GOMAXPROCS(0) - 1
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x) // why x = 0 ?
}
