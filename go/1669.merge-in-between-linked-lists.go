package leetcodesolution

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=1669 lang=golang
 *
 * [1669] Merge In Between Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	n, i := list1, 0
	var na, nb *ListNode
	for {
		if i+1 == a {
			na = n
		}
		if i == b {
			nb = n.Next
			break
		}
		i++
		n = n.Next
	}

	na.Next = list2
	n2 := list2
	for n2.Next != nil {
		n2 = n2.Next
	}
	n2.Next = nb
	return list1
}

// @lc code=end
