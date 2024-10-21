package binarytreelevelordertraversal

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
	lvl  int
	node *TreeNode
}

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}

	l := 0
	nodes := []node{{lvl: 0, node: root}}
	for len(nodes) > 0 {
		i := 0
		vs := make([]int, 0)
		for i < len(nodes) && nodes[i].lvl == l {
			if nodes[i].node.Left != nil {
				nodes = append(nodes, node{lvl: l + 1, node: nodes[i].node.Left})
			}
			if nodes[i].node.Right != nil {
				nodes = append(nodes, node{lvl: l + 1, node: nodes[i].node.Right})
			}

			vs = append(vs, nodes[i].node.Val)
			i++
		}
		ans = append(ans, vs)
		nodes = nodes[i:]
		l++
	}
	return
}

// @lc code=end
