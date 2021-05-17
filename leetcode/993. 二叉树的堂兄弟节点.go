package main

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

func isCousins(root *TreeNode, x int, y int) bool {
	var foo func(node *TreeNode, tar int, depth int, parent *TreeNode) (int, *TreeNode)
	foo = func(node *TreeNode, tar int, depth int, parent *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, parent
		}
		if node.Val == tar {
			return depth, parent
		}
		level, rPatent := foo(node.Left, tar, depth+1, node)
		if level == 0 {
			level, rPatent = foo(node.Right, tar, depth+1, node)
		}
		return level, rPatent
	}
	parent := TreeNode{}
	depth1, parent1 := foo(root, x, 0, &parent)
	depth2, parent2 := foo(root, y, 0, &parent)
	if depth1 == depth2 && parent1 != parent2 {
		return true
	}
	return false
}
