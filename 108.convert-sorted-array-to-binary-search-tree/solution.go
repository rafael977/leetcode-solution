package convertsortedarraytobinarysearchtree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	var build func(i, j int) *TreeNode
	build = func(i, j int) *TreeNode {
		if i > j {
			return nil
		}
		m := i + (j-i)/2
		root := &TreeNode{Val: nums[m]}
		root.Left = build(i, m-1)
		root.Right = build(m+1, j)

		return root
	}

	return build(0, len(nums)-1)
}

// @lc code=end
