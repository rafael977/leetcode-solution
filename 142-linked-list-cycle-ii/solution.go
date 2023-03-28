package linkedlistcycleii

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

// @lc code=end
