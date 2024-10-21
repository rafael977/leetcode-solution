package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2058 lang=golang
 *
 * [2058] Find the Minimum and Maximum Number of Nodes Between Critical Points
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	pos := make([]int, 0)
	p := 1
	for head != nil && head.Next != nil && head.Next.Next != nil {
		if (head.Next.Val > head.Val && head.Next.Val > head.Next.Next.Val) ||
			(head.Next.Val < head.Val && head.Next.Val < head.Next.Next.Val) {
			pos = append(pos, p)
		}
		p++
		head = head.Next
	}
	if len(pos) < 2 {
		return []int{-1, -1}
	}

	mind := pos[1] - pos[0]
	for i := 2; i < len(pos); i++ {
		mind = min(mind, pos[i]-pos[i-1])
	}
	return []int{mind, pos[len(pos)-1] - pos[0]}
}

// @lc code=end
