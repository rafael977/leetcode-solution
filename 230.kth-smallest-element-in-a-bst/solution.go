package kthsmallestelementinabst

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) (ans int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		k--
		if k == 0 {
			ans = root.Val
		}
		dfs(root.Right)
	}

	dfs(root)
	return
}

// @lc code=end
