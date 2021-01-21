package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	r := isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	return r
}

func main() {
	t1 := TreeNode{Val: 1}
	t2 := TreeNode{Val: 2}
	t3 := TreeNode{Val: 3}
	t1.Left = &t2
	t2.Left = &t3

	T1 := TreeNode{Val: 1}
	T2 := TreeNode{Val: 2}
	T3 := TreeNode{Val: 1}
	T1.Left = &T2
	T2.Left = &T3

	r := isSameTree(&t1, &t2)
	fmt.Println(r)
}
