package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	r, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
		}
		if v == 0 {
			if r < count {
				r = count
			}
			count = 0
		}
	}
	if r < count {
		r = count
	}
	return r
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	r := findMaxConsecutiveOnes(nums)
	fmt.Println(r)
}
