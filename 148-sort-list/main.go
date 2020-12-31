package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddle(head)
	r := mid.Next
	mid.Next = nil

	ll := sortList(head)
	lr := sortList(r)

	return merge(ll, lr)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func main() {
	head := BuildLinkedList("4,2,1,3")
	PrintLinkedList(sortList(head))

	head = BuildLinkedList("-1,5,3,4,0")
	PrintLinkedList(sortList(head))

	head = BuildLinkedList("")
	PrintLinkedList(sortList(head))
}
