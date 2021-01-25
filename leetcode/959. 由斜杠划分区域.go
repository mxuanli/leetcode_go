package main

import "fmt"

func regionsBySlashes(grid []string) int {
	n := len(grid)
	count := n * n * 4
	father := make(map[int]int)
	for i := 0; i < count; i++ {
		father[i] = i
	}
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
			count--
		}
	}
	var getId func(row, col, i int) int
	getId = func(row, col, i int) int {
		id := (row*n+col)*4 + i
		return id
	}
	for row := 0; row < n; row++ {
		line := grid[row]
		for col := 0; col < n; col++ {
			g := line[col]
			if row > 0 {
				x := getId(row, col, 0)
				y := getId(row-1, col, 2)
				union(x, y)
			}
			if col > 0 {
				x := getId(row, col, 3)
				y := getId(row, col-1, 1)
				union(x, y)
			}
			if g != '/' {
				x := getId(row, col, 0)
				y := getId(row, col, 1)
				union(x, y)
				x = getId(row, col, 2)
				y = getId(row, col, 3)
				union(x, y)
			}
			if g != '\\' {
				x := getId(row, col, 1)
				y := getId(row, col, 2)
				union(x, y)
				x = getId(row, col, 0)
				y = getId(row, col, 3)
				union(x, y)
			}
		}
	}
	return count
}

func main() {
	grid := []string{" /", "  "}
	r := regionsBySlashes(grid)
	fmt.Println(r)
}
