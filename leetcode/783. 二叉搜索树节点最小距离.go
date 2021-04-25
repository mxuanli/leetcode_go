/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 **/
package main

import "math"

func minDiffInBST(root *TreeNode) int {
	var tmp []int
	foo(root, &tmp)
	r := int(math.Pow10(5) + 1)
	for i := 1; i < len(tmp); i++ {
		if tmp[i]-tmp[i-1] < r {
			r = tmp[i] - tmp[i-1]
		}
	}
	return r
}

func foo(root *TreeNode, tmp *[]int) {
	if root == nil {
		return
	}
	foo(root.Left, tmp)
	*tmp = append(*tmp, root.Val)
	foo(root.Right, tmp)
}
