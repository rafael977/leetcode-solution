package deepestleavessum

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1302 lang=golang
 *
 * [1302] Deepest Leaves Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type QItem struct {
	*TreeNode
	Lvl int
}

func deepestLeavesSum(root *TreeNode) int {
	q := []QItem{{TreeNode: root, Lvl: 0}}

	sum := 0
	lvl := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if lvl == node.Lvl {
			sum += node.Val
		} else {
			sum = node.Val
			lvl = node.Lvl
		}

		if node.Left != nil {
			q = append(q, QItem{TreeNode: node.Left, Lvl: node.Lvl + 1})
		}
		if node.Right != nil {
			q = append(q, QItem{TreeNode: node.Right, Lvl: node.Lvl + 1})
		}
	}

	return sum
}

// @lc code=end
