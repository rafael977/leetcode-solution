package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1325 lang=golang
 *
 * [1325] Delete Leaves With a Given Value
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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(node, p *TreeNode, isLeft bool)
	dfs = func(node, p *TreeNode, isLeft bool) {
		if node.Left == nil && node.Right == nil {
			if node.Val == target {
				if isLeft {
					p.Left = nil
				} else {
					p.Right = nil
				}
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, node, true)
		}
		if node.Right != nil {
			dfs(node.Right, node, false)
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == target {
				if isLeft {
					p.Left = nil
				} else {
					p.Right = nil
				}
			}
		}
	}
	dummy := &TreeNode{Left: root}
	dfs(root, dummy, true)
	return dummy.Left
}

// @lc code=end
