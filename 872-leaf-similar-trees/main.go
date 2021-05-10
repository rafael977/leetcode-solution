package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1, a2 := bfs(root1), bfs(root2)
	if len(a1) != len(a2) {
		return false
	}

	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	r := make([]int, 0)
	r = append(r, bfs(root.Left)...)
	r = append(r, bfs(root.Right)...)
	return r
}

func main() {
	root1, root2 := BuildTree("3,5,1,6,2,9,8,null,null,7,4"), BuildTree("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8")
	fmt.Println(leafSimilar(root1, root2))

	root1, root2 = BuildTree("1"), BuildTree("1")
	fmt.Println(leafSimilar(root1, root2))

	root1, root2 = BuildTree("1"), BuildTree("2")
	fmt.Println(leafSimilar(root1, root2))

	root1, root2 = BuildTree("1,2"), BuildTree("2,2")
	fmt.Println(leafSimilar(root1, root2))
}
