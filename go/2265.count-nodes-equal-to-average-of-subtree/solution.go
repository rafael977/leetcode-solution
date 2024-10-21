package countnodesequaltoaverageofsubtree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2265 lang=golang
 *
 * [2265] Count Nodes Equal to Average of Subtree
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
func averageOfSubtree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(tn *TreeNode) (int, int) {
		if tn == nil {
			return 0, 0
		}
		ls, ln := dfs(tn.Left)
		rs, rn := dfs(tn.Right)
		sum := ls + tn.Val + rs
		num := ln + 1 + rn
		if sum/num == tn.Val {
			ans++
		}
		return sum, num
	}

	dfs(root)
	return
}

// @lc code=end
