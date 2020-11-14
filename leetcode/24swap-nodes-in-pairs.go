package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func swapPairs(head *ListNode1) *ListNode1 {
	return foo(head)
}

func foo(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	// 存储下一阶段的head
	tmp := head.Next.Next
	// 把next的next值设置位head，反转链表
	next := head.Next
	next.Next = head
	head.Next = foo(tmp) // 连接链表
	return next
}

func main() {
	var ln1 ListNode1
	var ln2 ListNode1
	var ln3 ListNode1
	var ln4 ListNode1
	ln1.Val = 1
	ln2.Val = 2
	ln3.Val = 3
	ln4.Val = 4
	ln1.Next = &ln2
	ln2.Next = &ln3
	ln3.Next = &ln4
	swapPairs(&ln1)
}
