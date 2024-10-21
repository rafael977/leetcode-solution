package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1530 lang=golang
 *
 * [1530] Number of Good Leaf Nodes Pairs
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
func countPairs(root *TreeNode, distance int) (ans int) {
	l := make([]string, 0)
	var dfs func(n *TreeNode, p string)
	dfs = func(n *TreeNode, p string) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			l = append(l, p)
			return
		}
		dfs(n.Left, p+"L")
		dfs(n.Right, p+"R")
	}
	dfs(root, "")

	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			k := 0
			for k < len(l[i]) && k < len(l[j]) && l[i][k] == l[j][k] {
				k++
			}
			if len(l[i])+len(l[j])-2*k <= distance {
				ans++
			}
		}
	}
	return
}

// @lc code=end
