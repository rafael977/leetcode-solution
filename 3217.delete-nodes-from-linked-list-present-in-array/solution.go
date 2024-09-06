package deletenodesfromlinkedlistpresentinarray

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=3217 lang=golang
 *
 * [3217] Delete Nodes From Linked List Present in Array
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	nm := make(map[int]bool)
	for _, n := range nums {
		nm[n] = true
	}
	tmp := &ListNode{Next: head}
	c := tmp
	for c.Next != nil {
		if nm[c.Next.Val] {
			c.Next = c.Next.Next
			continue
		}
		c = c.Next
	}

	return tmp.Next
}

// @lc code=end
