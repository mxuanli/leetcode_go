package main

import (
	"fmt"
	"sort"
)

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
	var edgesAddIndex [][]int
	// 给边添加下标
	for i, edge := range edges {
		edge = append(edge, i)
		edgesAddIndex = append(edgesAddIndex, edge)
	}
	// 排序
	sort.Slice(edgesAddIndex, func(i, j int) bool {
		return edgesAddIndex[i][2] < edgesAddIndex[j][2]
	})
	// 计算最小树
	total := 0
	num := 1
	for _, edge := range edgesAddIndex {
		x := edge[0]
		y := edge[1]
		w := edge[2]
		if isConnected(x, y) == false {
			total = total + w
			num = num + 1
			if num == n {
				break
			}
		}
	}
	// 遍历
	// 先连接当前边之后的总权值
	// 去掉当前边的权值
	// 和total比较，如果total和total0相等，说明是有效边
	// 如果total0和total1不相等，则是关键边
	// 否则是伪关键边
}

func main() {
	n := 5
	edges := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}
	r := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(r)
}
