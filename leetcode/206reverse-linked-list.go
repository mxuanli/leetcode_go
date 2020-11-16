package main

import "fmt"

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
type ListNode4 struct {
	Val  int
	Next *ListNode4
}

func reverseList2(head *ListNode4) *ListNode4 {
	/*
		tmp和head双指针，tmp在前，head在后，
		每次让head的next指向tmp，实现局部反转，然后以后tmp和head向后一位
	*/
	var tmp *ListNode4
	for head != nil {
		headNn := head.Next
		head.Next = tmp
		tmp = head
		head = headNn
	}
	return tmp
}

func reverseList(head *ListNode4) *ListNode4 {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head // 让后边的值的Next指向自己反转链表
	head.Next = nil       // 避免最后head.Next的Next指向head导致死循环
	return ret
}

func main() {
	var l1 ListNode4
	var l2 ListNode4
	var l3 ListNode4
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l1.Next = &l2
	l2.Next = &l3
	head := &l1
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	r := reverseList(&l1)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
