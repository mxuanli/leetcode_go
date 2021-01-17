package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var rSlice []string
	lenN := len(nums)
	if lenN == 0 {
		return rSlice
	}
	nums = append(nums, nums[0]-10)
	start, end := 0, 0
	for i := 0; i < lenN; i++ {
		if nums[i+1] == nums[i]+1 {
			end++
		} else {
			if start == end {
				r := strconv.Itoa(nums[start])
				rSlice = append(rSlice, r)
			} else {
				r := strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end])
				rSlice = append(rSlice, r)
			}
			start, end = i+1, i+1
		}
	}
	return rSlice
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	r := summaryRanges(nums)
	fmt.Println(r)
}
