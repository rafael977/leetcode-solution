package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return 1 + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	root := BuildTree("3,9,20,null,null,15,7")
	fmt.Println(maxDepth(root))

	root = BuildTree("1,null,2")
	fmt.Println(maxDepth(root))

	root = BuildTree("")
	fmt.Println(maxDepth(root))

	root = BuildTree("0")
	fmt.Println(maxDepth(root))
}
