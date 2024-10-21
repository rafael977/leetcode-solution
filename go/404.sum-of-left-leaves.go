package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, isLeft bool)
	dfs = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			ans += node.Val
		}
		dfs(node.Left, true)
		dfs(node.Right, false)
	}
	dfs(root, false)
	return
}

// @lc code=end
