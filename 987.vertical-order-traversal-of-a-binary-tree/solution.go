package verticalordertraversalofabinarytree

import (
	"sort"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/*
 * @lc app=leetcode id=987 lang=golang
 *
 * [987] Vertical Order Traversal of a Binary Tree
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

type pos struct {
	Val, Row, Col int
}

func verticalTraversal(root *TreeNode) (ans [][]int) {
	ps := make([]pos, 0)

	var dfs func(root *TreeNode, r, c int)
	dfs = func(root *TreeNode, r, c int) {
		if root == nil {
			return
		}
		ps = append(ps, pos{Val: root.Val, Row: r, Col: c})
		dfs(root.Left, r+1, c-1)
		dfs(root.Right, r+1, c+1)
	}

	dfs(root, 0, 0)
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].Col < ps[j].Col {
			return true
		} else if ps[i].Col > ps[j].Col {
			return false
		} else if ps[i].Row < ps[j].Row {
			return true
		} else if ps[i].Row > ps[j].Row {
			return false
		} else {
			return ps[i].Val < ps[j].Val
		}
	})

	initLvl := ps[0].Col
	lvls := ps[len(ps)-1].Col - initLvl + 1
	ans = make([][]int, lvls)
	for _, p := range ps {
		lvl := p.Col - initLvl
		if len(ans[lvl]) == 0 {
			ans[lvl] = make([]int, 0)
		}

		ans[lvl] = append(ans[lvl], p.Val)
	}
	return
}

// @lc code=end
