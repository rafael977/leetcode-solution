package binarytreepreordertraversal

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) (ans []int) {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return
}

// @lc code=end
