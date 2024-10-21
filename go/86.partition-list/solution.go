package partitionlist

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dl, dr := &ListNode{Val: 0}, &ListNode{Val: 0}
	lc, rc := dl, dr
	for head != nil {
		if head.Val < x {
			lc.Next = head
			lc = head
		} else {
			rc.Next = head
			rc = head
		}
		head = head.Next
	}

	lc.Next = dr.Next
	rc.Next = nil

	return dl.Next
}

// @lc code=end
