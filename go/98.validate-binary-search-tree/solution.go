package validatebinarysearchtree

import (
	"math"

	//lint:ignore ST1001 dot import
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	return isInRange(root, math.MinInt64, math.MaxInt64)
}

func isInRange(root *TreeNode, lo, hi int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lo || root.Val >= hi {
		return false
	}

	return isInRange(root.Left, lo, root.Val) && isInRange(root.Right, root.Val, hi)
}

// @lc code=end
