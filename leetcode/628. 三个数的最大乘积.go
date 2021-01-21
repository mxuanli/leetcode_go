package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums) - 1
	max1 := nums[0] * nums[1] * nums[n]
	max2 := nums[n] * nums[n-1] * nums[n-2]
	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

func main() {
	nums := []int{1, 2, 3}
	r := maximumProduct(nums)
	fmt.Println(r)
}
