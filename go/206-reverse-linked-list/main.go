package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, head := head, head.Next
	pre.Next = nil

	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre
}

func main() {
	head := BuildLinkedList("1,2,3,4,5")
	PrintLinkedList(reverseList(head))
}
