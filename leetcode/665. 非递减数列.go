package main

import "fmt"

func checkPossibility(nums []int) bool {
	count := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			count++
			if count > 1 {
				return false
			}
			if i == 1 || (i > 1 && nums[i-2] < nums[i]) {
				nums[i-1] = nums[i]
			}
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}

func main() {
	nums := []int{4, 2, 3}
	r := checkPossibility(nums)
	fmt.Println(r)
}
