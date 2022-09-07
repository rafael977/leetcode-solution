package constructstringfrombinarytree

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	lStr, rStr := tree2str(root.Left), tree2str(root.Right)
	if len(lStr) > 0 && len(rStr) > 0 {
		return fmt.Sprintf("%d(%s)(%s)", root.Val, lStr, rStr)
	} else if len(lStr) > 0 {
		return fmt.Sprintf("%d(%s)", root.Val, lStr)
	} else if len(rStr) > 0 {
		return fmt.Sprintf("%d()(%s)", root.Val, rStr)
	}
	return fmt.Sprintf("%d", root.Val)
}

// @lc code=end
