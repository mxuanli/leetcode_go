package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	s = s + "."
	left, right := 0, 0
	var r [][]int
	for ; right < len(s); right++ {
		if s[left] == s[right] {
			continue
		}
		if right-left >= 3 {
			L := []int{left, right - 1}
			r = append(r, L)
		}
		left = right
	}
	return r
}

func main() {
	s := "abcdddeeeeaabbbcd"
	r := largeGroupPositions(s)
	fmt.Println(r)
}
