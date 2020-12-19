package main

import "fmt"

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return nums[start]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	r := findMin(nums)
	fmt.Println(r)
}
