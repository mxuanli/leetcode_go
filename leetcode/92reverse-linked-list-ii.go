package main

import "fmt"

/**
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode6 struct {
	Val  int
	Next *ListNode6
}

func reverseBetween(head *ListNode6, m int, n int) *ListNode6 {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode6{Val: 0}
	tmp.Next = head
	var mid *ListNode6 // m的前一个节点，mid -> 反转链表 ->
	head = tmp
	for i := 0; i < m; i++ {
		mid = head
		head = head.Next
	}
	var pre *ListNode6 // 反转链表交换节点
	tmp2 := head       // 转换前的m节点，反转后用于连接链表的后半部分 mid -> 反转链表 -> 后
	j := m
	for head != nil && j <= n {
		nnl := head.Next
		head.Next = pre
		pre = head
		head = nnl
		j++
	}
	mid.Next = pre
	tmp2.Next = head
	return tmp.Next
}

func reverseBetween1(head *ListNode6, m int, n int) *ListNode6 {
	if m == 1 {
		return foo(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

var nnl *ListNode6

func foo(head *ListNode6, n int) *ListNode6 {
	if n == 1 {
		nnl = head.Next
		return head
	}
	last := foo(head.Next, n-1)
	head.Next.Next = head
	head.Next = nnl
	return last
}

func main() {
	var l1 ListNode6
	var l2 ListNode6
	var l3 ListNode6
	var l4 ListNode6
	var l5 ListNode6
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 4
	l5.Val = 5
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	head := &l1
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	m := 2
	n := 2
	fmt.Println("-----")
	r := reverseBetween(&l4, m, n)
	//r := foo(&l1, n)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
