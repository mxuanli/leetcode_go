package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	for start+1 < end {
		mid := (end + start) / 2 // 设置中间点
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { // 如果中间点大于目标值，则移动end
			end = mid
		} else if nums[mid] < target { // 如果中间点小于目标值，则移动start
			start = mid
		}
	}
	return -1
}

func main() {
	nums := []int{2, 5}
	target := 5
	r := search(nums, target)
	fmt.Println(r)
}
