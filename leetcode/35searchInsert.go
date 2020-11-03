package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		}
	}
	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] <= target {
		return len(nums)
	}
	return 0
}

func main() {
	nums := []int{1, 3, 4, 6}
	target := 2
	r := searchInsert(nums, target)
	fmt.Println(r)
}
