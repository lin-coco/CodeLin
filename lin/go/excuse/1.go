package main

import (
	"fmt"
	"time"
)

/*
https://github.com/lifei6671/interview-go/blob/master/question/q001.md
交替打印数字和字母
问题描述

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/
func main() {
	c1 := make(chan struct{}, 0)
	c2 := make(chan struct{}, 0)
	go func() {
		for i := 1; i <= 26; i++ {
			fmt.Print(i)
			c1 <- struct{}{}
			<-c2
		}
	}()

	go func() {
		for i := 'a'; i <= 'z'; i++ {
			<-c1
			fmt.Printf("%c", i)
			c2 <- struct{}{}
		}
	}()
	time.Sleep(time.Second)
}
