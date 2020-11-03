package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var list []int
	backtrack(0, list, &result, nums)
	return result
}

func backtrack(i int, list []int, result *[][]int, nums []int) {
	z := make([]int, len(list))
	copy(z, list)
	*result = append(*result, z)
	for j := i; j < len(nums); j++ {
		list = append(list, nums[j])
		backtrack(j+1, list, result, nums)
		list = list[0 : len(list)-1]
	}
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	c := subsets(nums)
	fmt.Println(c)
}
