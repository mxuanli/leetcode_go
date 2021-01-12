package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	father := make(map[int]int)
	for i := 0; i < len(s); i++ {
		father[i] = i
	}
	// 查找
	var find func(x int) int
	find = func(x int) int {
		root := x
		for father[root] != root {
			root = father[root]
		}
		// 压缩
		for x != root {
			fatherCode := father[x]
			father[x] = root
			x = fatherCode
		}
		return root
	}
	// 合并
	var merge func(x int, y int)
	merge = func(x int, y int) {
		rx := find(x)
		ry := find(y)
		if rx != ry {
			father[rx] = ry
		}
	}
	// 创建
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		merge(x, y)
	}
	nodeSlice := make([][]int, len(s))
	// 获取连通节点
	for i := 0; i < len(s); i++ {
		fatherNode := find(i)
		nodeSlice[fatherNode] = append(nodeSlice[fatherNode], i)
	}
	r := s
	for i := 0; i < len(nodeSlice); i++ {
		// 排序下标
		node := nodeSlice[i]
		sort.Ints(node)
		var byteSlice []string
		// 排序字符
		for _, j := range node {
			byteSlice = append(byteSlice, string(byte(s[j])))
		}
		sort.Strings(byteSlice)
		// 交换
		for k, v := range node {
			r = r[:v] + byteSlice[k] + r[v+1:]
		}
	}
	return r
}

func main() {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}, {0, 2}}
	r := smallestStringWithSwaps(s, pairs)
	fmt.Println(r)
}
