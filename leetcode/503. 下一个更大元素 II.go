package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i, _ := range res {
		res[i] = -1
	}
	var stack []int
	for i := 0; i < n*2; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return res
}

func main() {
	nums := []int{1, 2, 1}
	r := nextGreaterElements(nums)
	fmt.Println(r)
}
