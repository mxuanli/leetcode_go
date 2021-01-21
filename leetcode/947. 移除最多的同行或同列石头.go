package main

import (
	"fmt"
)

func removeStones(stones [][]int) int {
	father := make(map[int]int)
	var find func(x int) int
	find = func(x int) int {
		root := x
		for _, exists := father[root]; exists; _, exists = father[root] {
			root = father[root]
		}
		for root != x {
			fatherRoot := father[x]
			father[x] = root
			x = fatherRoot
		}
		return root
	}

	var union func(x int, y int)
	union = func(x int, y int) {
		xf := find(x)
		yf := find(y)
		if xf != yf {
			father[xf] = yf
		}
	}

	if len(stones) == 0 || len(stones) == 1 {
		return 0
	}
	for _, stone := range stones {
		x, y := stone[0], stone[1]+15000
		union(x, y)
	}
	rMap := make(map[int]int)
	for _, stone := range stones {
		x, _ := stone[0], stone[1]
		xv := find(x)
		rMap[xv] = 1
	}
	r := len(stones) - len(rMap)
	return r
}

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	r := removeStones(stones)
	fmt.Println(r)
}
