package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}

		l, r := 0, 0
		if node.Left != nil {
			l = dfs(node.Left) + 1
		}
		if node.Right != nil {
			r = dfs(node.Right) + 1
		}
		ans = max(ans, l+r)
		return max(l, r)
	}

	dfs(root)
	return
}

// @lc code=end
