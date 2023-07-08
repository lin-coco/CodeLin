package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	//fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	//fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	//fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	//fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	//fmt.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))
	part1 := Part1{}
	part2 := Part2{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}

type Part2 struct {
	e byte
	c int8
	a bool
	b int32
	d int64
}

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

func test3() {

}

func test2() {
	s := "a小b"
	d := s[1:]
	f := s[2:]
	fmt.Println(unsafe.Pointer(&s))
	fmt.Println(unsafe.Pointer(&d))
	fmt.Println(unsafe.Pointer(&f))
}

func test1() {
	ints := make([]int32, 0)
	ints = append(ints, 0)
	ints = append(ints, 1)
	//fmt.Println(len(ints), cap(ints))
	//for i := 1; i < 10; i++ {
	//	ints = append(ints, i)
	//	fmt.Println(len(ints), cap(ints))
	//}
	pointer := unsafe.Pointer(&ints[0])
	pointer1 := unsafe.Pointer(&ints[1])
	fmt.Println(pointer)
	fmt.Println(pointer1)
}
