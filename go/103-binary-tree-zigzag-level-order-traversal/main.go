package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	return helper([]*TreeNode{root}, res, true)
}

func helper(root []*TreeNode, res [][]int, dir bool) [][]int {
	lvl := make([]int, 0)
	next := make([]*TreeNode, 0)
	for _, v := range root {
		if v == nil {
			continue
		}
		lvl = append(lvl, v.Val)
		next = append(next, v.Left, v.Right)
	}
	if len(next) == 0 {
		return res
	}

	if !dir {
		reverse(lvl)
	}
	res = append(res, lvl)
	return helper(next, res, !dir)
}

func reverse(lvl []int) {
	for i, j := 0, len(lvl)-1; i < j; {
		lvl[i], lvl[j] = lvl[j], lvl[i]
		i++
		j--
	}
}

func main() {
	root := BuildTree("3,9,20,null,null,15,7")
	fmt.Printf("%v", zigzagLevelOrder(root))
}
