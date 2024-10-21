package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
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
func findBottomLeftValue(root *TreeNode) (ans int) {
	lvl := -1
	var dfs func(node *TreeNode, l int)
	dfs = func(node *TreeNode, l int) {
		if node.Left == nil && node.Right == nil {
			if l > lvl {
				ans = node.Val
				lvl = l
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, l+1)
		}
		if node.Right != nil {
			dfs(node.Right, l+1)
		}
	}

	dfs(root, 0)
	return
}

// @lc code=end
