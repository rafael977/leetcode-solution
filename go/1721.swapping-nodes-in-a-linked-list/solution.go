package swappingnodesinalinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1721 lang=golang
 *
 * [1721] Swapping Nodes in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	n := 0
	iter := head
	for iter != nil {
		n++
		iter = iter.Next
	}

	iter = head
	var i, j *ListNode
	cnt := 0
	for iter != nil {
		if cnt == k-1 {
			i = iter
		}
		if cnt == n-k {
			j = iter
		}

		if i != nil && j != nil {
			i.Val, j.Val = j.Val, i.Val
			break
		}
		cnt++
		iter = iter.Next
	}

	return head
}

// @lc code=end
