package main

import "fmt"

type TreeS struct {
	value interface{}
	left  *TreeS
	right *TreeS
}

func makeTree() TreeS {
	var t TreeS
	t.value = 10
	var t0 TreeS
	t0.value = 9
	var t1 TreeS
	t1.value = 20
	var t2 TreeS
	t2.value = 15
	var t3 TreeS
	t3.value = 35
	var t4 TreeS
	t4.value = 5
	var t5 TreeS
	t5.value = 7
	var t6 TreeS
	t6.value = 3
	t.left = &t5
	t5.left = &t4
	t5.right = &t0
	t4.left = &t6
	t.right = &t1
	t1.left = &t2
	t1.right = &t3
	return t
}

func preorderTraversal(root *TreeS) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

func preorderMidTraversal(root *TreeS) {
	if root == nil {
		return
	}
	preorderMidTraversal(root.left)
	fmt.Println(root.value)
	preorderMidTraversal(root.right)
}

func main() {
	t := makeTree()
	preorderMidTraversal(&t)
}
