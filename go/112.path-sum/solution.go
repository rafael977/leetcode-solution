package pathsum

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return root.Left != nil && hasPathSum(root.Left, targetSum-root.Val) ||
		root.Right != nil && hasPathSum(root.Right, targetSum-root.Val)
}

// @lc code=end
