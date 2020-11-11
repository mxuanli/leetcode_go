package main

import "fmt"

func fib(N int) int {
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}

func fib2(N int) int {
	// 递归
	if N == 0 {
		return N
	}
	if N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

func main() {
	N := 10
	r := fib2(N)
	fmt.Println(r)
}
