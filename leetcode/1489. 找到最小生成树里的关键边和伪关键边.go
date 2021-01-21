package main

import "fmt"

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	father := make(map[int]int)
	var find func(x int) int
	// 查找
	find = func(x int) int {
		root := x
		// 查找
		for father[root] != root {
			root = father[root]
		}
		// 压缩
		if root != x {
			fatherRoot := father[x]
			father[x] = root
			x = fatherRoot
		}
		return root
	}
	// 合并
	var union func(x int, y int)
	union = func(x int, y int) {
		xf := find(x)
		yf := find(y)
		if xf != yf {
			father[xf] = yf
		}
	}
	// 判断连通
	var isConnected func(int, int) bool
	isConnected = func(x int, y int) bool {
		xf := find(x)
		yf := find(y)
		if xf != yf {
			return false
		}
		return true
	}
	// 给边添加下标
	for i, edge := range edges {
		edge = append(edge, i)
	}
	fmt.Println(edges)
	return [][]int{}
}

func main() {
	n := 5
	edges := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}
	r := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(r)
}
