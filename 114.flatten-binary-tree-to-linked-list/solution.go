package flattenbinarytreetolinkedlist

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	list := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	c := root
	for i := 1; i < len(list); i++ {
		c.Left = nil
		c.Right = list[i]
		c = c.Right
	}
}

// @lc code=end
