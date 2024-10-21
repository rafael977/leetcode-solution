package pseudopalindromicpathsinabinarytree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1457 lang=golang
 *
 * [1457] Pseudo-Palindromic Paths in a Binary Tree
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
func pseudoPalindromicPaths(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode, path [9]int)
	dfs = func(root *TreeNode, path [9]int) {
		path[root.Val-1]++
		if root.Left == nil && root.Right == nil {
			if isPseudoPalindromic(path) {
				ans++
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}
	dfs(root, [9]int{})
	return
}

func isPseudoPalindromic(path [9]int) bool {
	oddCount := 0
	for _, v := range path {
		if v%2 == 1 {
			oddCount++
		}
	}
	return oddCount <= 1
}

// @lc code=end
