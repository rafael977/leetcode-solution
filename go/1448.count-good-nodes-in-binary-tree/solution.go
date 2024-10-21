package countgoodnodesinbinarytree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) (ans int) {
	var check func(*TreeNode, int)
	check = func(root *TreeNode, maxVal int) {
		if root == nil {
			return
		}
		if root.Val >= maxVal {
			ans++
		}
		maxVal = max(root.Val, maxVal)
		check(root.Left, maxVal)
		check(root.Right, maxVal)
	}

	check(root, (-10e4 - 1))
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
