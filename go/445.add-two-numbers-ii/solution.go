package addtwonumbersii

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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
	s1, s2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var c *ListNode
	o := 0
	for len(s1) > 0 && len(s2) > 0 {
		n1, n2 := s1[len(s1)-1], s2[len(s2)-1]
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]

		sum := n1 + n2 + o
		c = &ListNode{Val: sum % 10, Next: c}
		o = sum / 10
	}

	for len(s1) > 0 {
		n1 := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		sum := n1 + o
		c = &ListNode{Val: sum % 10, Next: c}
		o = sum / 10
	}
	for len(s2) > 0 {
		n2 := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]

		sum := n2 + o
		c = &ListNode{Val: sum % 10, Next: c}
		o = sum / 10
	}
	if o != 0 {
		c = &ListNode{Val: o, Next: c}
	}

	return c
}

// @lc code=end
