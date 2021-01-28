package main

import "fmt"

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
		ufs.count = ufs.count - 1
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

func maxNumEdgesToRemove2(n int, edges [][]int) int {
	r := 0
	ufsA := UnionFindSet{map[int]int{}, n}
	ufsB := UnionFindSet{map[int]int{}, n}
	var slice1, slice2, slice3 [][]int
	for _, edge := range edges {
		x, y := edge[1], edge[2]
		ufsA.father[x] = x
		ufsA.father[y] = y
		ufsB.father[x] = x
		ufsB.father[y] = y
	}
	for _, edge := range edges {
		t := edge[0]
		if t == 3 {
			slice3 = append(slice3, edge)
		} else if t == 2 {
			slice2 = append(slice2, edge)
		} else if t == 1 {
			slice1 = append(slice1, edge)
		}
	}
	// 公共边
	for _, edge := range slice3 {
		x, y := edge[1], edge[2]
		a := ufsA.isConnected(x, y)
		b := ufsB.isConnected(x, y)
		if a == false {
			ufsA.union(x, y)
		}
		if b == false {
			ufsB.union(x, y)
		}
		if a == true && b == true {
			r++
		}
	}
	// A边
	for _, edge := range slice1 {
		x, y := edge[1], edge[2]
		a := ufsA.isConnected(x, y)
		if a == false {
			ufsA.union(x, y)
		} else {
			r++
		}
	}
	// B边
	for _, edge := range slice2 {
		x, y := edge[1], edge[2]
		b := ufsB.isConnected(x, y)
		if b == false {
			ufsB.union(x, y)
		} else {
			r++
		}
	}
	if ufsA.count > 1 || ufsB.count > 1 {
		return -1
	}
	return r
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	r := 0
	ufsA := UnionFindSet{map[int]int{}, n}
	ufsB := UnionFindSet{map[int]int{}, n}
	for _, edge := range edges {
		x, y := edge[1], edge[2]
		ufsA.father[x] = x
		ufsA.father[y] = y
		ufsB.father[x] = x
		ufsB.father[y] = y
	}
	// 公共边
	for _, edge := range edges {
		t := edge[0]
		if t == 3 {
			x, y := edge[1], edge[2]
			a := ufsA.isConnected(x, y)
			b := ufsB.isConnected(x, y)
			if a == false {
				ufsA.union(x, y)
			}
			if b == false {
				ufsB.union(x, y)
			}
			if a == true && b == true {
				r++
			}
		}
	}
	// A、B边
	for _, edge := range edges {
		t := edge[0]
		if t == 1 {
			x, y := edge[1], edge[2]
			a := ufsA.isConnected(x, y)
			if a == false {
				ufsA.union(x, y)
			} else {
				r++
			}
		}
		if t == 2 {
			x, y := edge[1], edge[2]
			b := ufsB.isConnected(x, y)
			if b == false {
				ufsB.union(x, y)
			} else {
				r++
			}
		}
	}
	if ufsA.count > 1 || ufsB.count > 1 {
		return -1
	}
	return r
}

func main() {
	n := 4
	edges := [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}
	r := maxNumEdgesToRemove(n, edges)
	fmt.Println(r)
}
