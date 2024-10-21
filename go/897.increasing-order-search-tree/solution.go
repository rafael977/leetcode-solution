package increasingordersearchtree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=897 lang=golang
 *
 * [897] Increasing Order Search Tree
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
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		cur.Right = &TreeNode{Val: root.Val}
		cur = cur.Right
		dfs(root.Right)
	}

	dfs(root)
	return dummy.Right
}

// @lc code=end
