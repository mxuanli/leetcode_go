package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	r := nums[0] + nums[n-1]
	var max func(a int, b int) int
	max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n/2; i++ {
		r = max(r, nums[i]+nums[n-i-1])
	}
	return r
}
