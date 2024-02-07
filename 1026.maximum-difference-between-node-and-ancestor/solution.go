package maximumdifferencebetweennodeandancestor

import (
	"math"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
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
func maxAncestorDiff(root *TreeNode) (ans int) {
	an := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		for _, v := range an {
			ans = max(ans, int(math.Abs(float64(node.Val-v))))
		}
		an = append(an, node.Val)
		dfs(node.Left)
		dfs(node.Right)
		an = an[:len(an)-1]
	}

	dfs(root)
	return
}

// @lc code=end
