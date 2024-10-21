package main

//lint:ignore ST1001 dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}
	prev := tmp

	for prev.Next != nil {
		c := prev.Next
		n := prev.Next.Next
		for n != nil && n.Val == c.Val {
			n = n.Next
		}

		if n != c.Next {
			prev.Next = n
		} else {
			prev = prev.Next
		}
	}

	return tmp.Next
}

// @lc code=end
