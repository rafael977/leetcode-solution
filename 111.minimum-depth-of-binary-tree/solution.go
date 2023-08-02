package minimumdepthofbinarytree

import (
	"math"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

type NodeLevel struct {
	node *TreeNode
	lvl  int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MaxInt
	q := []NodeLevel{{node: root, lvl: 1}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.node.Left == nil && n.node.Right == nil {
			ans = min(ans, n.lvl)
			continue
		}

		if n.node.Left != nil {
			q = append(q, NodeLevel{node: n.node.Left, lvl: n.lvl + 1})
		}
		if n.node.Right != nil {
			q = append(q, NodeLevel{node: n.node.Right, lvl: n.lvl + 1})
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
