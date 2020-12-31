package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	var p, c *ListNode = nil, head
	i := 1
	for i < m {
		p = c
		c = c.Next
		i++
	}

	x, y := p, c

	for i <= n {
		tmp := c.Next
		c.Next = p
		p = c
		c = tmp
		i++
	}

	y.Next = c
	if x != nil {
		x.Next = p
		return head
	} else {
		return p
	}
}

func main() {
	head := BuildLinkedList("1,2,3,4,5")
	PrintLinkedList(reverseBetween(head, 2, 4))

	head = BuildLinkedList("1,2,3,4,5")
	PrintLinkedList(reverseBetween(head, 1, 4))

	head = BuildLinkedList("1,2,3,4,5")
	PrintLinkedList(reverseBetween(head, 1, 5))
}
