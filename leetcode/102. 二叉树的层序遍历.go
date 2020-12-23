package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	if root == nil {
		return r
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var tmp []int
		lenQ := len(q)
		for i := 0; i < lenQ; i++ {
			t := q[0]
			tmp = append(tmp, t.Val)
			q = q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		r = append(r, tmp)
	}
	return r
}

func main() {
	var root TreeNode
	root.Val = 3
	var T1 TreeNode
	T1.Val = 9
	var T2 TreeNode
	T2.Val = 20
	var T3 TreeNode
	T3.Val = 15
	var T4 TreeNode
	T4.Val = 7
	root.Left = &T1
	root.Right = &T2
	T2.Left = &T3
	T2.Right = &T4
	r := levelOrder(&root)
	fmt.Println(r)
}
