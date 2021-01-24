package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	start, end := 0, 1
	maxN := 0
	for end < n {
		for end < n && nums[end-1] < nums[end] {
			end++
		}
		if maxN < (end - start) {
			maxN = end - start
		}
		start = end
		end++
	}
	return maxN
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	r := findLengthOfLCIS(nums)
	fmt.Println(r)
}
