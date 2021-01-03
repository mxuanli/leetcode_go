package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		sort.Ints(stones)
		y := stones[len(stones)-1]
		x := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if x != y {
			stones = append(stones, y-x)
		}
	}
	if len(stones) == 0 {
		return 0
	} else {
		return stones[0]
	}
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	r := lastStoneWeight(stones)
	fmt.Println(r)
}
