package main

import "fmt"

func fib1(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

func main() {
	N := 4
	r := fib(N)
	fmt.Println(r)
}
