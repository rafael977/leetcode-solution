package main

//lint:ignore ST1001 dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}
	cur := tmp
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		n1, n2 := cur.Next, cur.Next.Next
		n1.Next = n2.Next
		n2.Next = n1
		cur.Next = n2
		cur = n1
	}

	return tmp.Next
}

// @lc code=end
