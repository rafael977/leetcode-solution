package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: 0, Next: head}
	cur, next := dummy, head
	for next != nil && next.Next != nil {
		if next.Val != next.Next.Val {
			cur.Next = next
			cur = cur.Next
			next = cur.Next
			continue
		}
		for next != nil && next.Next != nil && next.Val == next.Next.Val {
			next = next.Next
		}
		cur.Next = next.Next
		next = next.Next
	}

	return dummy.Next
}

func main() {
	head := BuildLinkedList("1,2,3,3,4,4,5")
	head = deleteDuplicates(head)
	PrintLinkedList(head)

	head = BuildLinkedList("1,1,1,2,3")
	head = deleteDuplicates(head)
	PrintLinkedList(head)

	head = BuildLinkedList("1,2,3,3,4,4,5,5")
	head = deleteDuplicates(head)
	PrintLinkedList(head)

	head = BuildLinkedList("1,1,2,2,3,3,4,4,5,5")
	head = deleteDuplicates(head)
	PrintLinkedList(head)
}
