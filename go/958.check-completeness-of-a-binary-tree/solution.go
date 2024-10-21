package checkcompletenessofabinarytree

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
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
func isCompleteTree(root *TreeNode) bool {
	n := 0
	maxId := 0
	var dfs func(root *TreeNode, id int)
	dfs = func(root *TreeNode, id int) {
		if root == nil {
			return
		}
		n++
		if id > maxId {
			maxId = id
		}
		dfs(root.Left, id*2)
		dfs(root.Right, id*2+1)
	}
	dfs(root, 1)

	return n == maxId
}

// @lc code=end
