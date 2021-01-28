package main

import "fmt"

func sum(nums []int) int {
	n := 0
	for _, num := range nums {
		n += num
	}
	return n
}

func pivotIndex(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		sum1 := sum(nums[:i])
		sum2 := sum(nums[i+1:])
		if sum1 == sum2 {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	r := pivotIndex(nums)
	fmt.Println(r)
}
