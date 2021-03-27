package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	tmp := head
	n := 1
	for tmp.Next != nil {
		tmp = tmp.Next
		n += 1
	}
	tmp.Next = head
	k %= n
	for i := 0; i < n-1-k; i++ {
		head = head.Next
	}
	newHead := head.Next
	head.Next = nil
	return newHead
}
