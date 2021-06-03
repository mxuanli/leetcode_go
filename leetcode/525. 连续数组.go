package main

import "fmt"

func findMaxLength(nums []int) int {
	pre := make(map[int]int)
	pre[0] = -1
	tmp := 0
	r := 0
	for i, num := range nums {
		if num != 0 {
			tmp += num
		} else {
			tmp -= 1
		}
		if index, ok := pre[tmp]; ok {
			diff := i - index
			if diff >= 2 && diff > r {
				r = diff
			}
		} else {
			pre[tmp] = i
		}
	}
	return r
}

func main() {
	nums := []int{0, 0, 1, 0, 0, 0, 1, 1}
	r := findMaxLength(nums)
	fmt.Println(r)
}
