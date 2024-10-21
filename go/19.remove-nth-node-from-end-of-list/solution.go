package removenthnodefromendoflist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}
	prev, slow, fast := tmp, head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next
	return tmp.Next
}

// @lc code=end
