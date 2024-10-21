package main

//lint:ignore ST1001 dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	tail := head
	l := 1
	for tail.Next != nil {
		tail = tail.Next
		l++
	}
	tail.Next = head
	for i := 0; i < l-k%l; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil

	return head
}

// @lc code=end
