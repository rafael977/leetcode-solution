package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	f, s := head.Next, head
	for f != nil && f.Next != nil {
		if f == s {
			return true
		}
		s = s.Next
		f = f.Next.Next
	}

	return false
}

func main() {
	head := BuildLinkedList("3,2,0,-4")
	head = GenCycle(head, 1)
	fmt.Printf("%v\n", hasCycle(head))

	head = BuildLinkedList("1,2")
	head = GenCycle(head, 0)
	fmt.Printf("%v\n", hasCycle(head))

	head = BuildLinkedList("1,2")
	head = GenCycle(head, -1)
	fmt.Printf("%v\n", hasCycle(head))
}
