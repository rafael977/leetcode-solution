package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	tail, next := head, head.Next
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	head.Next = tail.Next
	head.Next.Next = next
	tail.Next = nil

	reorderList(next)
}

func main() {
	head := BuildLinkedList("1,2,3,4")
	reorderList(head)
	PrintLinkedList(head)

	head = BuildLinkedList("1,2,3,4,5")
	reorderList(head)
	PrintLinkedList(head)
}
