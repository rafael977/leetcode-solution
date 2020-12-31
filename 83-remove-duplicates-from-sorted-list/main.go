package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	v := head.Val
	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val == v {
			prev.Next = cur.Next
		} else {
			v = cur.Val
			prev = cur
		}
		cur = cur.Next
	}

	return head
}

func main() {
	head := BuildLinkedList("1,1,2")
	PrintLinkedList(deleteDuplicates(head))

	head = BuildLinkedList("1,1,2,3,3")
	PrintLinkedList(deleteDuplicates(head))
}
