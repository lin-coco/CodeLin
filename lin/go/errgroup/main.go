package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

/*
怎么说呢？
比较无用吧
我想用多线程就是要并发执行的，这给搞成串行了
*/

func main() {
	group := new(errgroup.Group)

	nums := []int{-1, 0, 1}
	for _, num := range nums {
		num := num
		group.Go(func() error {
			res, err := output(num)
			fmt.Println(res)
			return err
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}
}

func output(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("math: square root error!")
	}
	return num, nil
}
