package binarytreepruning

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=814 lang=golang
 *
 * [814] Binary Tree Pruning
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
func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		l1, r1 := dfs(root.Left), dfs(root.Right)
		if !l1 {
			root.Left = nil
		}
		if !r1 {
			root.Right = nil
		}
		return root.Val == 1 || l1 || r1
	}
	if !dfs(root) {
		return nil
	}

	return root
}

// @lc code=end
