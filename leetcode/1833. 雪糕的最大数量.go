package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	r := 0
	if costs[0] > coins {
		return r
	}
	for _, c := range costs {
		if c <= coins {
			coins -= c
			r += 1
		}
	}
	return r
}

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7
	r := maxIceCream(costs, coins)
	fmt.Println(r)
}
