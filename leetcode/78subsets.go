package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var list []int
	foo(nums, 0, list, &result)
	return result
}

func foo(nums []int, j int, list []int, result *[][]int) {
	l := make([]int, len(list))
	copy(l, list)
	*result = append(*result, l)
	for i := j; i < len(nums); i++ {
		list = append(list, nums[i])
		foo(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	r := subsets(nums)
	fmt.Println(r)
}
