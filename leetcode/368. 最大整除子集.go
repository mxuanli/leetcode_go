package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	dp := make([]int, n)
	g := make([]int, n)
	for i := 0; i < n; i++ {
		length := 1
		pre := 0
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if dp[j]+1 > length {
					length = dp[j] + 1
					pre = j
				}
			}
		}
		dp[i] = length
		g[i] = pre
	}
	MaxLength := -1
	index := 0
	for i, v := range dp {
		if v > MaxLength {
			MaxLength = v
			index = i
		}
	}
	var res []int
	for len(res) < MaxLength {
		res = append(res, nums[index])
		index = g[index]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
