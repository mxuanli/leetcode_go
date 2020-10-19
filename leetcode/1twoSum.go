package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var resultSlice []int
	var mapA map[int]int
	mapA = make(map[int]int)
	for i, j := range nums {
		diff := target - j
		diffM, ok := mapA[diff]
		if ok {
			return []int{i, diffM}
		} else {
			mapA[j] = i
		}
	}
	return resultSlice
}

func twoSum2(nums []int, target int) []int {
	var resultSlice []int
	var mapA map[int]int
	mapA = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		diffM, ok := mapA[diff]
		if ok {
			return []int{i, diffM}
		}
		mapA[nums[i]] = i
	}
	return resultSlice
}

func main() {
	a := []int{2, 7, 11, 15}
	target := 9
	resultSlice := twoSum2(a, target)
	fmt.Println(resultSlice)
}
