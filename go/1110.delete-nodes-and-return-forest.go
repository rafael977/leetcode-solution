package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1110 lang=golang
 *
 * [1110] Delete Nodes And Return Forest
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
func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
	tdm := make(map[int]bool)
	for _, v := range to_delete {
		tdm[v] = true
	}
	var dfs func(n *TreeNode) *TreeNode
	dfs = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}
		n.Left = dfs(n.Left)
		n.Right = dfs(n.Right)
		if !tdm[n.Val] {
			return n
		}
		if n.Left != nil {
			ans = append(ans, n.Left)
		}
		if n.Right != nil {
			ans = append(ans, n.Right)
		}
		return nil
	}
	if n := dfs(root); n != nil {
		ans = append(ans, n)
	}
	return
}

// @lc code=end
