package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi := 0
	for si := 0; si < len(s); si++ {
		if s[si] >= g[gi] {
			gi++
		}
		if gi == len(g) {
			return gi
		}
	}
	return gi
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	r := findContentChildren(g, s)
	fmt.Println(r)
}
