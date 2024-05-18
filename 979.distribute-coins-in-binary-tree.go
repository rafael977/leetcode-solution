package leetcodesolution

import (
	"math"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=979 lang=golang
 *
 * [979] Distribute Coins in Binary Tree
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
func distributeCoins(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		cl, nl := dfs(node.Left)
		cr, nr := dfs(node.Right)
		ans += int(math.Abs(float64(cl-nl)) + math.Abs(float64(cr-nr)))
		return cl + cr + node.Val, nl + nr + 1
	}
	dfs(root)
	return
}

// @lc code=end
