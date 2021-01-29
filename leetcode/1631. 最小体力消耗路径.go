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

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return 0 - x
}

func minimumEffortPath(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])
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
				upId := thisId - m
				tmp := [3]int{upId, thisId, abs(heights[i][j] - heights[i-1][j])}
				tmpSlice = append(tmpSlice, tmp)
			}
			if j > 0 {
				leftId := thisId - 1
				tmp := [3]int{leftId, thisId, abs(heights[i][j] - heights[i][j-1])}
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
	heights := [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	r := minimumEffortPath(heights)
	fmt.Println(r)
}
