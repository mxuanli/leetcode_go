package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {
	nums := x ^ y
	r := 0
	for ; nums > 0; nums >>= 1 {
		r += nums & 1
	}
	return r
}

func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	x := 4
	y := 5
	r := hammingDistance(x, y)
	fmt.Println(r)
}
