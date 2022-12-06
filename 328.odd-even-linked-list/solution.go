package oddevenlinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd, even := head, head.Next
	oc, ec := odd, even
	i := 0
	c := head.Next.Next
	for c != nil {
		if i == 0 {
			oc.Next = c
			oc = oc.Next
		} else {
			ec.Next = c
			ec = ec.Next
		}
		i ^= 1
		c = c.Next
	}
	oc.Next = even
	ec.Next = nil
	return odd
}

// @lc code=end
