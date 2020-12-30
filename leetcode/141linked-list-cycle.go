package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode5 struct {
	Val  int
	Next *ListNode5
}

func hasCycle(head *ListNode5) bool {
	if head == nil {
		return false
	}
	left := head
	right := head.Next
	for right != nil && right.Next != nil {
		if left == right {
			return true
		}
		left = left.Next
		right = right.Next.Next
	}
	return false
}

func hasCycle1(head *ListNode5) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}
	tmp := head.Next
	head.Next = head
	return hasCycle1(tmp)
}

func main() {
	var l1 ListNode5
	var l2 ListNode5
	var l3 ListNode5
	l1.Val = 1
	l2.Val = 1
	l3.Val = 2
	l1.Next = &l2
	l2.Next = &l1
	//head := &l1
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}
	r := hasCycle1(&l1)
	fmt.Println(r)
}
