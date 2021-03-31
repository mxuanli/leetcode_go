package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var path []int
	var result [][]int
	sort.Ints(nums)
	foo(&result, path, 0, nums)
	return result
}

func foo(result *[][]int, path []int, index int, nums []int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		foo(result, path, i+1, nums)
		path = path[0 : len(path)-1]
	}
}
