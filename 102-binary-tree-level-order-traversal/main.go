package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/pkg/structure"
)

func levelOrder(root *TreeNode) [][]int {
	return helper([]*TreeNode{root})
}

func helper(nodes []*TreeNode) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	next := make([]*TreeNode, 0)
	for _, v := range nodes {
		if v == nil {
			continue
		}
		cur = append(cur, v.Val)
		next = append(next, v.Left, v.Right)
	}

	if len(cur) == 0 {
		return res
	}
	res = append(res, cur)
	res = append(res, helper(next)...)

	return res
}

func main() {
	root := BuildTree("3,9,20,null,null,15,7")
	fmt.Printf("%#v", levelOrder(root))
}
