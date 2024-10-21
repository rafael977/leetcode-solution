package binarytreerightsideview

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
type node struct {
	lvl int
	n   *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)

	q := []node{{lvl: 0, n: root}}

	l := 0
	for len(q) > 0 {
		i := 0
		for i < len(q) && q[i].lvl == l {
			if q[i].n.Left != nil {
				q = append(q, node{lvl: l + 1, n: q[i].n.Left})
			}
			if q[i].n.Right != nil {
				q = append(q, node{lvl: l + 1, n: q[i].n.Right})
			}
			i++
		}
		ans = append(ans, q[i-1].n.Val)
		q = q[i:]
		l++
	}

	return ans
}

// @lc code=end
