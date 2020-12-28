package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/pkg/structure"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	arr := inorderTraversal(root.Left)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)

	return arr
}

func main() {
	root := BuildTree("1,null,2,3")
	PrintTree(root)

	fmt.Println(inorderTraversal(root))
}
