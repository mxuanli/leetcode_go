package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	r := 0
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		r += root.Val
	}
	if root.Val > low {
		r += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		r += rangeSumBST(root.Right, low, high)
	}
	return r
}
