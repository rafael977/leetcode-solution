package main

import (
	"fmt"
	"math"

	. "github.com/rafael977/leetcode-solution/pkg/tree"
)

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

func main() {
	root := BuildTree("2,1,3")
	fmt.Println(isValidBST(root))

	root = BuildTree("5,1,4,null,null,3,6")
	fmt.Println(isValidBST(root))

	root = BuildTree("1,1")
	fmt.Println(isValidBST(root))
}
