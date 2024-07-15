package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=2196 lang=golang
 *
 * [2196] Create Binary Tree From Descriptions
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
func createBinaryTree(descriptions [][]int) *TreeNode {
	tn := make(map[int]*TreeNode)
	p := make(map[int]int)
	for _, d := range descriptions {
		n, c := tn[d[0]], tn[d[1]]
		if n == nil {
			n = &TreeNode{Val: d[0]}
			tn[d[0]] = n
		}
		if c == nil {
			c = &TreeNode{Val: d[1]}
			tn[d[1]] = c
		}
		if d[2] == 1 {
			n.Left = c
		} else if d[2] == 0 {
			n.Right = c
		}

		if _, ok := p[d[0]]; !ok {
			p[d[0]] = 0
		}
		p[d[1]]++
	}
	for k, v := range p {
		if v == 0 {
			return tn[k]
		}
	}
	return nil
}

// @lc code=end
