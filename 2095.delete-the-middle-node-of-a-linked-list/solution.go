package deletethemiddlenodeofalinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2095 lang=golang
 *
 * [2095] Delete the Middle Node of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next
	return dummy.Next
}

// @lc code=end
