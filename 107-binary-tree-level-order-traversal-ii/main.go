package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	return helper([]*TreeNode{root}, res)
}

func helper(root []*TreeNode, res [][]int) [][]int {
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

	res = append(helper(next, res), lvl)
	return res
}

func main() {
	root := BuildTree("3,9,20,null,null,15,7")
	fmt.Printf("%v", levelOrderBottom(root))
}
