package reverselinkedlistii

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
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

// @lc code=end
