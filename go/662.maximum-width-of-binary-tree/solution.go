package maximumwidthofbinarytree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
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

type item struct {
	node *TreeNode
	idx  int
}

func widthOfBinaryTree(root *TreeNode) (ans int) {
	ans = 1
	arr := []item{{root, 1}}
	for len(arr) > 0 {
		next := make([]item, 0)
		for _, v := range arr {
			if v.node.Left != nil {
				next = append(next, item{v.node.Left, 2 * v.idx})
			}
			if v.node.Right != nil {
				next = append(next, item{v.node.Right, 2*v.idx + 1})
			}
		}
		if len(next) > 0 {
			ans = max(ans, next[len(next)-1].idx-next[0].idx+1)
		}
		arr = next
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
