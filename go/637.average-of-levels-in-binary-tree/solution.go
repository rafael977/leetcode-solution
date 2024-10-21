package averageoflevelsinbinarytree

import . "github.com/rafael977/leetcode-solution/datastruct"

/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}

	lvl := []*TreeNode{root}
	for len(lvl) > 0 {
		newLvl := make([]*TreeNode, 0)
		sum := 0
		for _, n := range lvl {
			sum += n.Val
			if n.Left != nil {
				newLvl = append(newLvl, n.Left)
			}
			if n.Right != nil {
				newLvl = append(newLvl, n.Right)
			}
		}

		ans = append(ans, float64(sum)/float64(len(lvl)))
		lvl = newLvl
	}
	return
}

// @lc code=end
