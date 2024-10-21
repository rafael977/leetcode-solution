package recoverbinarysearchtree

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	nodes := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		nodes = append(nodes, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	i, j := -1, -1
	for k := 0; k < len(nodes)-1; k++ {
		if nodes[k] > nodes[k+1] {
			if i == -1 {
				i = k
			}
			j = k + 1
		}
	}

	cnt := 2
	var recover func(root *TreeNode, x, y int)
	recover = func(root *TreeNode, x, y int) {
		if root == nil {
			return
		}

		if root.Val == x {
			root.Val = y
			cnt--
		} else if root.Val == y {
			root.Val = x
			cnt--
		}
		if cnt == 0 {
			return
		}
		recover(root.Left, x, y)
		recover(root.Right, x, y)
	}

	recover(root, nodes[i], nodes[j])
}

// @lc code=end
