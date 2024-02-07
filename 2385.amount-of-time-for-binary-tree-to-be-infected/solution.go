package amountoftimeforbinarytreetobeinfected

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2385 lang=golang
 *
 * [2385] Amount of Time for Binary Tree to Be Infected
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
func amountOfTime(root *TreeNode, start int) int {
	g := make(map[int][]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			if _, ok := g[node.Val]; !ok {
				g[node.Val] = make([]int, 0)
			}
			g[node.Val] = append(g[node.Val], node.Left.Val)
			if _, ok := g[node.Left.Val]; !ok {
				g[node.Left.Val] = make([]int, 0)
			}
			g[node.Left.Val] = append(g[node.Left.Val], node.Val)
			dfs(node.Left)
		}
		if node.Right != nil {
			if _, ok := g[node.Val]; !ok {
				g[node.Val] = make([]int, 0)
			}
			g[node.Val] = append(g[node.Val], node.Right.Val)
			if _, ok := g[node.Right.Val]; !ok {
				g[node.Right.Val] = make([]int, 0)
			}
			g[node.Right.Val] = append(g[node.Right.Val], node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)

	visited := make(map[int]bool)
	visited[start] = true
	q := []int{start}
	ans := -1
	for len(q) > 0 {
		ans++
		nq := make([]int, 0)
		for _, v := range q {
			for _, vv := range g[v] {
				if !visited[vv] {
					visited[vv] = true
					nq = append(nq, vv)
				}
			}
		}
		q = nq
	}
	return ans
}

// @lc code=end
