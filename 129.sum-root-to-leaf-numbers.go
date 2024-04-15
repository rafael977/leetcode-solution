package leetcodesolution

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */

import . "github.com/rafael977/leetcode-solution/datastruct"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, ps int)
	dfs = func(node *TreeNode, ps int) {
		if node == nil {
			return
		}

		sum := 10*ps + node.Val
		if node.Left == nil && node.Right == nil {
			ans += sum
			return
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, 0)
	return
}

// @lc code=end
