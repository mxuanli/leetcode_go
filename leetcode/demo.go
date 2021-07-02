package main

import "fmt"

func foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	n := foo()
	for i := 0; i < 10; i++ {
		fmt.Println(n())
	}

}
