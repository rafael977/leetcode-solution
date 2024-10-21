package leetcodesolution

/*
 * @lc app=leetcode id=988 lang=golang
 *
 * [988] Smallest String Starting From Leaf
 */

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) (ans string) {
	var dfs func(node *TreeNode, p string)
	dfs = func(node *TreeNode, p string) {
		if node == nil {
			return
		}
		s := fmt.Sprintf("%c%s", 'a'+node.Val, p)
		if node.Left == nil && node.Right == nil && (ans == "" || s < ans) {
			ans = s
			return
		}

		dfs(node.Left, s)
		dfs(node.Right, s)
	}

	dfs(root, "")
	return
}

// @lc code=end
