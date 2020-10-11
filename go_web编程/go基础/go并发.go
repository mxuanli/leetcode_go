package main

import (
	"fmt"
	"runtime"
)

func SayHi(a string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(a)
	}
}

func main() {
	go SayHi("hello")
	SayHi("world")
}
