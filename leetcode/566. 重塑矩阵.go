package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = nums[i/n][i%n]
	}
	return res
}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	r := 1
	c := 4
	res := matrixReshape(nums, r, c)
	fmt.Println(res)
}
