package main

import (
	"fmt"
	"sort"
)

func numEquivDominoPairs(dominoes [][]int) int {
	tmp := make(map[[2]int]int)
	r := 0
	numsSlice := [2]int{}
	for _, nums := range dominoes {
		sort.Ints(nums)
		x, y := nums[0], nums[1]
		if x > y {
			numsSlice = [2]int{x, y}
		} else {
			numsSlice = [2]int{y, x}
		}
		r = r + tmp[numsSlice]
		tmp[numsSlice]++
	}
	return r
}

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	r := numEquivDominoPairs(dominoes)
	fmt.Println(r)
}
