package twosumivinputisabst

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if m[root.Val] {
			return true
		}
		m[k-root.Val] = true
		return dfs(root.Left) || dfs(root.Right)
	}

	return dfs(root)
}

// @lc code=end
