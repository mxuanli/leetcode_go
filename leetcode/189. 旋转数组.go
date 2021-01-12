package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	lenN := len(nums)
	n1 := nums[lenN-k:]
	n2 := nums[0 : lenN-k]
	nums = append(n1, n2...)
	copy(nums, nums)
}

func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	fmt.Println(nums)
}

func main() {
	nums := []int{1, 2}
	k := 3
	rotate(nums, k)
}
