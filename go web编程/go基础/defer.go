package main

import (
	"fmt"
)

func main() {
	a := testDefer()
	fmt.Println(a)
}

func testDefer() int {
	fmt.Println(" world")
	defer fmt.Println("hello")
	fmt.Println("golang")
	return 0
}
