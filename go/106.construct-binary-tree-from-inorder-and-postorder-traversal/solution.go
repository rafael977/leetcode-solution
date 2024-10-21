package constructbinarytreefrominorderandpostordertraversal

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	v := postorder[n-1]
	i := 0
	for ; i < n; i++ {
		if inorder[i] == v {
			break
		}
	}
	return &TreeNode{
		Val:   v,
		Left:  buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:n-1]),
	}
}

// @lc code=end
