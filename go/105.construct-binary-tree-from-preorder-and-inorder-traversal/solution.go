package constructbinarytreefrompreorderandinordertraversal

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	idx := make(map[int]int)
	for i, v := range inorder {
		idx[v] = i
	}

	var build func(pol, poh, iol, ioh int) *TreeNode
	build = func(pol, poh, iol, ioh int) *TreeNode {
		if pol > poh {
			return nil
		}

		root := &TreeNode{Val: preorder[pol]}
		ioi := idx[root.Val]
		lenLTree := ioi - iol
		root.Left = build(pol+1, pol+lenLTree, iol, ioi-1)
		root.Right = build(pol+lenLTree+1, poh, ioi+1, ioh)

		return root
	}

	return build(0, len(preorder)-1, 0, len(preorder)-1)
}

// @lc code=end
