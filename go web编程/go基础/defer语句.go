package main

import (
	"fmt"
)

func main() {
	a := testDefer()
	fmt.Println(a)
}

func testDefer() int {
	//执行到return，程序会逆序执行defer语句
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("golang")
	return 0
}
