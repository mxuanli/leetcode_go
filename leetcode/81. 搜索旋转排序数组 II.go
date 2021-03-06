package main

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		for start+1 < end && nums[start] == nums[start+1] {
			start++
		}
		for start+1 < end && nums[end] == nums[end-1] {
			end--
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[end] {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
		if nums[mid] > nums[end] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		}
	}
	if target == nums[start] || target == nums[end] {
		return true
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 7, 8, 1, 2}
	target := 0
	r := search(nums, target)
	fmt.Println(r)
}
