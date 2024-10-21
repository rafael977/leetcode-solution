package main

import (
	"fmt"
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	dfs(root, low, high, &sum)
	return sum
}

func dfs(root *TreeNode, low, high int, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Left, low, high, sum)
	if root.Val >= low && root.Val <= high {
		*sum += root.Val
	}
	dfs(root.Right, low, high, sum)
}

func main() {
	root := BuildTree("10,5,15,3,7,null,18")
	fmt.Println(rangeSumBST(root, 7, 15))
}
