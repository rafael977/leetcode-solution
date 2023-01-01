package rangesumofbst

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val == low {
		return root.Val + rangeSumBST(root.Right, low, high)
	}
	if root.Val == high {
		return root.Val + rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, root.Val) + rangeSumBST(root.Right, root.Val, high)
	}
}

// @lc code=end
