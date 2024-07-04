package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2181 lang=golang
 *
 * [2181] Merge Nodes in Between Zeros
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	p, c := head, head.Next

	sum := 0
	for c != nil {
		if c.Val == 0 {
			n := &ListNode{Val: sum}
			p.Next = n
			p = n
			c = c.Next
			sum = 0
			continue
		}

		sum += c.Val
		c = c.Next
	}

	return head.Next
}

// @lc code=end
