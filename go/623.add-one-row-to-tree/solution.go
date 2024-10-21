package addonerowtotree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	if depth == 2 {
		root.Left = &TreeNode{Val: val, Left: root.Left}
		root.Right = &TreeNode{Val: val, Right: root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}

// @lc code=end
