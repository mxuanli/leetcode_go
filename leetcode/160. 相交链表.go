package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	HA := headA
	HB := headB
	for HA != HB {
		if HA != nil {
			HA = HA.Next
		} else {
			HA = headB
		}
		if HB != nil {
			HB = HB.Next
		} else {
			HB = headA
		}
	}
	return HA
}
