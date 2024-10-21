package maximumtwinsumofalinkedlist

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2130 lang=golang
 *
 * [2130] Maximum Twin Sum of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) (ans int) {
	st := make([]int, 0)
	f, s := head, head
	for s != nil && s.Next != nil {
		st = append(st, f.Val)
		f = f.Next
		s = s.Next.Next
	}

	for f != nil {
		sum := st[len(st)-1] + f.Val
		ans = max(ans, sum)
		st = st[:len(st)-1]
		f = f.Next
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
