package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dl, dr := &ListNode{Val: 0}, &ListNode{Val: 0}
	lc, rc := dl, dr
	for head != nil {
		if head.Val < x {
			lc.Next = head
			lc = head
		} else {
			rc.Next = head
			rc = head
		}
		head = head.Next
	}

	lc.Next = dr.Next
	rc.Next = nil

	return dl.Next
}

func main() {
	head := BuildLinkedList("1,4,3,2,5,2")
	PrintLinkedList(partition(head, 3))

	head = BuildLinkedList("1,4,3,2,5,2")
	PrintLinkedList(partition(head, 1))

	head = BuildLinkedList("1,4,3,2,5,2")
	PrintLinkedList(partition(head, 5))
}
