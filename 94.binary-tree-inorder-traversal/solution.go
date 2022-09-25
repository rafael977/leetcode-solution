package binarytreeinordertraversal

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [97] Binary Tree Inorder Traversal
 */

// @lc code=start
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	arr := inorderTraversal(root.Left)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)

	return arr
}

// @lc code=end
