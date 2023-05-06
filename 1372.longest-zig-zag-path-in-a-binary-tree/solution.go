package longestzigzagpathinabinarytree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1372 lang=golang
 *
 * [1372] Longest ZigZag Path in a Binary Tree
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
func longestZigZag(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int, int)
	dfs = func(tn *TreeNode, l, r int) {
		ans = max(ans, max(l, r))
		if tn.Left != nil {
			dfs(tn.Left, r+1, 0)
		}
		if tn.Right != nil {
			dfs(tn.Right, 0, l+1)
		}
	}

	dfs(root, 0, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
