package main

//lint:ignore ST1001 dot import
import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return 1 + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// @lc code=end
