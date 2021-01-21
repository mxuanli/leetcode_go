package main

import (
	"fmt"
	"sort"
)

func minCostConnectPoints(points [][]int) int {
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
	// 绝对值
	var abs func(x int) int
	abs = func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	// 计算曼哈顿距离
	var manhattan func(x []int, y []int) int
	manhattan = func(x []int, y []int) int {
		r1 := abs(x[0] - y[0])
		r2 := abs(x[1] - y[1])
		return r1 + r2
	}
	// 初始化
	n := len(points)
	for i := 0; i < n; i++ {
		father[i] = i
	}
	// 计算所有边长
	var manhattanSlice [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := points[i]
			y := points[j]
			manhattanSlice = append(manhattanSlice, []int{manhattan(x, y), i, j})
		}
	}
	sort.Slice(manhattanSlice, func(i, j int) bool {
		return manhattanSlice[i][0] < manhattanSlice[j][0]
	})
	// 计算最短路径
	r := 0
	i := 1
	for _, manhattanNum := range manhattanSlice {
		length := manhattanNum[0]
		x := manhattanNum[1]
		y := manhattanNum[2]
		// 如果不在一个集合内就合并
		if isConnected(x, y) == false {
			union(x, y)
			r = r + length
			i += 1
			if i == n {
				break
			}
		}
	}
	return r
}

func main() {
	points := [][]int{{11, -6}, {9, -19}, {16, -13}, {4, -9}, {20, 4}, {20, 7}, {-9, 18}, {10, -15}, {-15, 3}, {6, 6}}
	r := minCostConnectPoints(points)
	fmt.Println(r)
}
