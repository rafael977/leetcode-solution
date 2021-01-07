package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func binaryTreePaths(root *TreeNode) []string {
	return helper(root, "")
}

func helper(root *TreeNode, path string) []string {
	if root == nil {
		return make([]string, 0)
	}

	if path == "" {
		path = fmt.Sprintf("%d", root.Val)
	} else {
		path = fmt.Sprintf("%s->%d", path, root.Val)
	}

	if root.Left == nil && root.Right == nil {
		return []string{path}
	}

	r := make([]string, 0)
	r = append(r, helper(root.Left, path)...)
	r = append(r, helper(root.Right, path)...)

	return r
}

func main() {
	root := BuildTree("1,2,3,null,5")
	fmt.Printf("%v\n", binaryTreePaths(root))
}
