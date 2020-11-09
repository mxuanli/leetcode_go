/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

请找出其中最小的元素。
*/

package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[end]
}

func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
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
	nums := []int{3, 4, 5, 6, 7, 1, 2}
	r := findMin2(nums)
	fmt.Println(r)
}
