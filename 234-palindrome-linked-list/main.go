package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	tail = reverse(tail)

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	f, s := head.Next, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	return s
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p, c *ListNode = nil, head
	for c != nil {
		tmp := c.Next
		c.Next = p
		p = c
		c = tmp
	}

	return p
}

func main() {
	head := BuildLinkedList("1,2")
	fmt.Printf("%v\n", isPalindrome(head))

	head = BuildLinkedList("1,2,2,1")
	fmt.Printf("%v\n", isPalindrome(head))

	head = BuildLinkedList("1,0,1")
	fmt.Printf("%v\n", isPalindrome(head))
}
