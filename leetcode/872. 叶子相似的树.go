package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var rootSlice1 []int
	var rootSlice2 []int

	var foo func(root *TreeNode, rootSlice *[]int)
	foo = func(root *TreeNode, rootSlice *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*rootSlice = append(*rootSlice, root.Val)
		}
		foo(root.Left, rootSlice)
		foo(root.Right, rootSlice)
	}
	foo(root1, &rootSlice1)
	foo(root2, &rootSlice2)
	if len(rootSlice1) != len(rootSlice2) {
		return false
	}
	for i, v := range rootSlice1 {
		if v != rootSlice2[i] {
			return false
		}
	}
	return true
}
