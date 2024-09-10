package insertgreatestcommondivisorsinlinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2807 lang=golang
 *
 * [2807] Insert Greatest Common Divisors in Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b

		}
		return b
	}

	cur, next := head, head.Next
	for cur != nil && next != nil {
		node := &ListNode{Val: gcd(cur.Val, next.Val), Next: next}
		cur.Next = node
		cur = next
		next = next.Next
	}
	return head
}

// @lc code=end
