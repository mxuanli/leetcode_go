package main

import fm "fmt"

func main() {
	var x int
	var y int
	x = 10
	x++
	y *= 10
	fm.Println(x)
	fm.Println(y)
	x = gcd(10, 7)
	fm.Println(x)
	f := fib(10)
	fm.Println(f)
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {

	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
