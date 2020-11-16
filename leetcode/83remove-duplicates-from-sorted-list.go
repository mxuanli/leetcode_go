package main

import "fmt"

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
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

func deleteDuplicates(head *ListNode1) *ListNode1 {
	tmp := head
	for tmp != nil {
		for tmp.Next != nil && tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return head
}

func main() {
	var l1 ListNode1
	var l2 ListNode1
	var l3 ListNode1
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
