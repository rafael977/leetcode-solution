package middleofthelinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=876 lang=golang
 *
 * [876] Middle of the Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	return s
}

// @lc code=end
