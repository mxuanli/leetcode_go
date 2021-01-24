package main

import "fmt"

func makeConnected(n int, connections [][]int) int {
	if n-1 > len(connections) {
		return -1
	}
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
			n = n - 1
		}
	}
	for _, connection := range connections {
		x, y := connection[0], connection[1]
		father[x] = x
		father[y] = y
	}
	for _, connection := range connections {
		x, y := connection[0], connection[1]
		union(x, y)
	}
	return n - 1
}

func main() {
	n := 4
	connections := [][]int{{0, 1}, {0, 2}, {1, 2}}
	r := makeConnected(n, connections)
	fmt.Println(r)
}
