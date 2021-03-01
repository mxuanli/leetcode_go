package main

import "fmt"

func findShortestSubArray(nums []int) int {
	map1 := make(map[int][3]int)
	for i, num := range nums {
		_, ok := map1[num]
		if ok == false {
			map1[num] = [3]int{1, i, i}
		} else {
			c := map1[num][0] + 1
			x := map1[num][1]
			y := i
			map1[num] = [3]int{c, x, y}
		}
	}
	count, r := 0, 50001
	for _, num := range map1 {
		c, x, y := num[0], num[1], num[2]
		if c > count {
			r = 50001
			count = c
			if y-x+1 < r {
				r = y - x + 1
			}
		} else if c == count {
			count = c
			if y-x+1 < r {
				r = y - x + 1
			}
		}
	}
	return r
}

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	r := findShortestSubArray(nums)
	fmt.Println(r)
}
