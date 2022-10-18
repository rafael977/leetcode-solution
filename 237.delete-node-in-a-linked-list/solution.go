package deletenodeinalinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}

// @lc code=end
