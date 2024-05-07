package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2816 lang=golang
 *
 * [2816] Double a Number Represented as a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	st := make([]*ListNode, 0)
	for head != nil {
		st = append(st, head)
		head = head.Next
	}

	tens := 0
	var cur *ListNode
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		cur = node
		double := node.Val*2 + tens
		node.Val = double % 10
		tens = double / 10
	}
	if tens == 1 {
		cur = &ListNode{Val: 1, Next: cur}
	}
	return cur
}

// @lc code=end
