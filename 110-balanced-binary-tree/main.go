package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/pkg/structure"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	b := false
	if abs(ld-rd) <= 1 {
		b = true
	}

	return b && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
	fmt.Println(isBalanced(root))

	root = BuildTree("1,2,2,3,3,null,null,4,4")
	fmt.Println(isBalanced(root))

	root = BuildTree("")
	fmt.Println(isBalanced(root))
}
