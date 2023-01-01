package leafsimilartrees

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1, leaf2 := dfs(root1), dfs(root2)

	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	leaf := dfs(root.Left)
	leaf = append(leaf, dfs(root.Right)...)
	return leaf
}

// @lc code=end
