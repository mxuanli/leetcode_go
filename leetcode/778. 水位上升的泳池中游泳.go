package main

import (
	"fmt"
	"sort"
)

type UnionFindSet struct {
	father map[int]int
}

func (ufs *UnionFindSet) find(x int) int {
	// 查找
	root := x
	for ufs.father[root] != root {
		root = ufs.father[root]
	}
	// 压缩
	if root != x {
		fatherRoot := ufs.father[x]
		ufs.father[x] = root
		x = fatherRoot
	}
	return root
}

func (ufs *UnionFindSet) union(x, y int) {
	// 合并
	xf := ufs.find(x)
	yf := ufs.find(y)
	if xf != yf {
		ufs.father[xf] = yf
	}
}

func (ufs *UnionFindSet) isConnected(x, y int) bool {
	// 判断连通
	xf := ufs.find(x)
	yf := ufs.find(y)
	if xf != yf {
		return false
	}
	return true
}

func max(x, y int) int {
	// 比大小
	if x > y {
		return x
	}
	return y
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	var tmpSlice [][3]int
	start, end := 0, n*m-1
	ufs := UnionFindSet{map[int]int{}}
	for i := 0; i <= end; i++ {
		ufs.father[i] = i
	}
	// 计算距离
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			thisId := i*m + j
			if i > 0 {
				downId := thisId - m
				tmp := [3]int{downId, thisId, max(grid[i][j], grid[i-1][j])}
				tmpSlice = append(tmpSlice, tmp)
			}
			if j > 0 {
				rightId := thisId - 1
				tmp := [3]int{rightId, thisId, max(grid[i][j], grid[i][j-1])}
				tmpSlice = append(tmpSlice, tmp)
			}
		}
	}
	// 排序
	sort.Slice(tmpSlice, func(i, j int) bool {
		return tmpSlice[i][2] < tmpSlice[j][2]
	})
	for _, tmp := range tmpSlice {
		x, y, v := tmp[0], tmp[1], tmp[2]
		ufs.union(x, y)
		if ufs.isConnected(start, end) {
			return v
		}
	}
	return 0
}

func main() {
	grid := [][]int{{0, 2}, {1, 3}}
	r := swimInWater(grid)
	fmt.Println(r)
}
