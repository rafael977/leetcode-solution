package symmetrictree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}

		return a.Val == b.Val && dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
	}

	return dfs(root.Left, root.Right)
}

// @lc code=end
