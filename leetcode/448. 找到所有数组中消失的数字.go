package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		x = 0 - x
	}
	return x
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] = nums[abs(v)-1] * -1
		}
	}
	r := []int{}
	for i, v := range nums {
		if v > 0 {
			r = append(r, i+1)
		}
	}
	return r
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	r := findDisappearedNumbers(nums)
	fmt.Println(r)
}
