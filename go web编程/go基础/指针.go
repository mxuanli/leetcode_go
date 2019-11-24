package main

import (
	"fmt"
)

func main() {
	pointerFunc()
}

func pointerFunc() {
	a := 10
	fmt.Println(*&a)
}
