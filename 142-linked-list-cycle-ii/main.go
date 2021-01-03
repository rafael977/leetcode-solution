package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for f != nil {
		if f.Next == nil {
			return nil
		}
		f = f.Next.Next
		s = s.Next

		if f == s {
			f = head
			for f != s {
				f = f.Next
				s = s.Next
			}
			return f
		}
	}
	return nil
}

func main() {
	head := BuildLinkedList("3,2,0,-4")
	head = GenCycle(head, 1)
	fmt.Printf("%v\n", detectCycle(head))

	head = BuildLinkedList("1,2")
	head = GenCycle(head, 0)
	fmt.Printf("%v\n", detectCycle(head))

	head = BuildLinkedList("1,2")
	head = GenCycle(head, -1)
	fmt.Printf("%v\n", detectCycle(head))
}
