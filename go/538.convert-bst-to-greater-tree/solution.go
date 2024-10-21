package convertbsttogreatertree

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}

// @lc code=end
