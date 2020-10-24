package main

import "fmt"

func reverse(x int) int {
	tmp := 0
	if x == 0 {
		return 0
	}
	for x != 0 {
		tmp = tmp * 10
		r := x % 10
		tmp = tmp + r
		x = x / 10
	}
	if tmp > (1<<31) || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func main() {
	x := 123
	tmp := reverse(x)
	fmt.Println(tmp)
}
