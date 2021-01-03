package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 分为小的和大的，再合并
	small := &ListNode{Val: 0}
	smallPre := small
	big := &ListNode{Val: 0}
	bigPre := big
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	small.Next = bigPre.Next
	big.Next = nil
	return smallPre.Next
}

func main() {
	a := ListNode{Val: 1}
	b := ListNode{Val: 4}
	c := ListNode{Val: 3}
	d := ListNode{Val: 2}
	e := ListNode{Val: 5}
	f := ListNode{Val: 2}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	x := 3
	r := partition(&a, x)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
