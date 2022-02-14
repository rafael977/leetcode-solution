package main

//lint:ignore ST1001 use dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{Val: 0, Next: nil}
	cur := h
	ten := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + ten
		ten = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + ten
		ten = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + ten
		ten = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		l2 = l2.Next
	}
	if ten != 0 {
		cur.Next = &ListNode{Val: ten}
	}

	return h.Next
}

// @lc code=end
