package leetcodesolution

/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	dummy := &TreeNode{Left: root}
	var dfs func(node *TreeNode, l int)
	dfs = func(node *TreeNode, l int) {
		if node == nil {
			return
		}
		if l == depth-1 {
			left, right := node.Left, node.Right
			node.Left = &TreeNode{Val: val, Left: left}
			node.Right = &TreeNode{Val: val, Right: right}
			return
		}
		dfs(node.Left, l+1)
		dfs(node.Right, l+1)
	}
	dfs(dummy, 0)
	return dummy.Left
}

// @lc code=end
