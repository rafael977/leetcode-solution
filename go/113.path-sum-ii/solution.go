package pathsumii

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var dfs func(root *TreeNode, path []int, sum int)
	dfs = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		sum += root.Val
		if root.Left == nil && root.Right == nil && sum == targetSum {
			ans = append(ans, append([]int{}, path...))
			return
		}

		if root.Left != nil {
			dfs(root.Left, path, sum)
		}
		if root.Right != nil {
			dfs(root.Right, path, sum)
		}
	}
	dfs(root, []int{}, 0)
	return
}

// @lc code=end
