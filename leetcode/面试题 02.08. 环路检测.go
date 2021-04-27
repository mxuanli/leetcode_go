package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil
}
