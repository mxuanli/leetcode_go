package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteDuplicates(head *ListNode2) *ListNode2 {
	if head == nil {
		return head
	}
	// head节点可能会因为重复被删除，所以新建一个节点放在最前边
	var tmp ListNode2
	tmp.Val = 0
	tmp.Next = head
	head = &tmp
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			// 使用一个临时变量保存重复的值，和此值相等的节点都删除
			val := head.Next.Val
			for head.Next != nil && head.Next.Val == val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return tmp.Next
}

func main() {
	var l1 ListNode2
	var l2 ListNode2
	var l3 ListNode2
	l1.Val = 1
	l2.Val = 1
	l3.Val = 2
	l1.Next = &l2
	l2.Next = &l3
	head := &l1
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	r := deleteDuplicates(&l1)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
