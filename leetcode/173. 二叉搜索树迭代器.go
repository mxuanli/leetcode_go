package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []int
}

func foo(root *TreeNode, stack *[]int) {
	if root == nil {
		return
	}
	foo(root.Left, stack)
	*stack = append(*stack, root.Val)
	foo(root.Right, stack)
	return
}

func Constructor(root *TreeNode) BSTIterator {
	var stack []int
	foo(root, &stack)
	b := BSTIterator{stack: stack}
	return b
}

func (this *BSTIterator) Next() int {
	r := this.stack[0]
	this.stack = this.stack[1:]
	return r
}

func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	}
	return true
}

func main() {
	var root TreeNode
	root.Val = 10
	b := Constructor(&root)
	a := b.Next()
	fmt.Println(a)
}
