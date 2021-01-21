package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	father := make(map[int]int)
	for i := 0; i < len(edges); i++ {
		for _, j := range edges[i] {
			father[j] = j
		}
	}
	// 查询
	var find func(x int) int
	find = func(x int) int {
		root := x
		for father[root] != root {
			root = father[root]
		}
		return root
	}
	// 合并
	var merge func(x int, y int)
	merge = func(x int, y int) {
		father[x] = y
	}
	// 遍历
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		xf, yf := find(x), find(y)
		if xf != yf {
			merge(xf, yf)
		} else {
			return []int{x, y}
		}
	}
	return []int{}
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	r := findRedundantConnection(edges)
	fmt.Println(r)
}
