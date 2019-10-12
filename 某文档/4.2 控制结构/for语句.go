package main

import (
	"fmt"
)

func main() {
	forFunc()
	for2Func()
	for3Func()
}

func forFunc() {
	for i := 0; i < 10; i++ {
		fmt.Println("for")
	}
}

func for2Func() {
	i := 0
	for i < 10 {
		i += 1
		fmt.Println("for2")
	}
}

func for3Func() {
	a := [...]string{"a", "b", "c"}
	for i, j := range a {
		fmt.Println(i, j)
	}
}
