package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1382 lang=golang
 *
 * [1382] Balance a Binary Search Tree
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
func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		arr = append(arr, n.Val)
		dfs(n.Right)
	}
	dfs(root)

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		node := &TreeNode{Val: arr[mid]}
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)
		return node
	}

	return build(0, len(arr)-1)
}

// @lc code=end
