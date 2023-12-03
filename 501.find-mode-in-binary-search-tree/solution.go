package findmodeinbinarysearchtree

import (
	"sort"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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
func findMode(root *TreeNode) (ans []int) {
	cnt := make(map[int]int)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		cnt[root.Val]++
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	max := 0
	for _, v := range cnt {
		if v > max {
			max = v
		}
	}
	for k, v := range cnt {
		if v == max {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return
}

// @lc code=end
