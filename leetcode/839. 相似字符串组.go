package main

import (
	"fmt"
)

type UnionFindSet struct {
	father map[int]int
	count  int
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
		ufs.count--
	}
}

func isSimilar(A, B string) bool {
	// 是否相似
	if A == B {
		return true
	}
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			count++
		}
	}
	if count <= 2 {
		return true
	}
	return false
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	ufs := UnionFindSet{map[int]int{}, n}
	for i := 0; i < n; i++ {
		ufs.father[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			A := strs[i]
			B := strs[j]
			if isSimilar(A, B) {
				ufs.union(i, j)
			}
		}
	}
	return ufs.count
}

func main() {
	strs := []string{"tars", "rats", "arts", "star"}
	r := numSimilarGroups(strs)
	fmt.Println(r)
}
