package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		// 跳过重复的值
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid := start + (end-start)/2
		// 如果中间点大于右边的点，那么最小值一定在右半边
		if nums[mid] >= nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}

func main() {
	nums := []int{3, 3, 3, 1, 3}
	a := findMin(nums)
	fmt.Println(a)
}
