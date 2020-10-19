package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1 *ListNode
	var n2 *ListNode
	n1, n2 = l1, l2
	carry := 0 // 进位
	var p int  // 存和
	for true {
		p = n1.Val + n2.Val + carry
		if p >= 10 {
			carry = p / 10
			p = p % 10
		} else {
			carry = 0
		}
		n1.Val = p
		if n1.Next != nil || n2.Next != nil || carry != 0 {
			if n1.Next == nil {
				n1.Next = &ListNode{Val: 0}
			}
			if n2.Next == nil {
				n2.Next = &ListNode{Val: 0}
			}
			n1 = n1.Next
			n2 = n2.Next
		} else {
			break
		}
	}
	return l1
}

func main() {
	l11 := ListNode{Val: 9}
	l12 := ListNode{Val: 9}
	l13 := ListNode{Val: 9}
	l14 := ListNode{Val: 9}
	l15 := ListNode{Val: 9}
	l16 := ListNode{Val: 9}
	l17 := ListNode{Val: 9}
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14
	l14.Next = &l15
	l15.Next = &l16
	l16.Next = &l17
	l21 := ListNode{Val: 9}
	l22 := ListNode{Val: 9}
	l23 := ListNode{Val: 9}
	l24 := ListNode{Val: 9}
	l21.Next = &l22
	l22.Next = &l23
	l23.Next = &l24
	n1 := addTwoNumbers(&l11, &l21)
	for n1 != nil {
		fmt.Println(n1.Val)
		n1 = n1.Next
	}
}
