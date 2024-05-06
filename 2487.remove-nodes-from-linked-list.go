package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2487 lang=golang
 *
 * [2487] Remove Nodes From Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	st := make([]*ListNode, 0)
	for head != nil {
		for len(st) > 0 && st[len(st)-1].Val < head.Val {
			st = st[:len(st)-1]
		}

		st = append(st, head)
		head = head.Next
	}
	dummyHead := &ListNode{}
	c := dummyHead
	for _, v := range st {
		c.Next = v
		c = c.Next
	}
	return dummyHead.Next
}

// @lc code=end
