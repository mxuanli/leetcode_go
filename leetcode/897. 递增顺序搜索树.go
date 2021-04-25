package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var nums []int
	foo(root, &nums)
	r := &TreeNode{Val: 0}
	tmp := r
	for _, num := range nums {
		node := TreeNode{Val: num}
		tmp.Right = &node
		tmp = tmp.Right
	}
	return r.Right
}

func foo(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	foo(root.Left, nums)
	*nums = append(*nums, root.Val)
	foo(root.Right, nums)
	return
}
