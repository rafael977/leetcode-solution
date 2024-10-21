package leetcodesolution

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=1609 lang=golang
 *
 * [1609] Even Odd Tree
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
func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	lvl := 0
	for len(q) > 0 {
		nq := make([]*TreeNode, 0)
		for i := range q {
			if q[i].Val%2 == lvl%2 ||
				(i+1 < len(q) &&
					((lvl%2 == 0 && q[i].Val >= q[i+1].Val) ||
						(lvl%2 == 1 && q[i].Val <= q[i+1].Val))) {
				return false
			}

			if q[i].Left != nil {
				nq = append(nq, q[i].Left)
			}
			if q[i].Right != nil {
				nq = append(nq, q[i].Right)
			}
		}
		lvl++
		q = nq
	}
	return true
}

// @lc code=end
