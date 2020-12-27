package main

import (
	"strconv"
	"strings"
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

// TreeNode defines for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := processInput("1,null,2,3")
	root := buildTree(arr, 0)

}

func processInput(input string) []string {
	arr := strings.Split(input, ",")
	for i := range arr {
		arr[i] = strings.TrimSpace(arr[i])
	}

	return arr
}

func buildTree(arr []string, idx int) *TreeNode {
	var root *TreeNode

	if idx < len(arr) {
		if arr[idx] == "null" {
			return nil
		}
		v, _ := strconv.Atoi(arr[idx])
		root = &TreeNode{v, nil, nil}
		root.Left = buildTree(arr, 2*idx+1)
		root.Right = buildTree(arr, 2*idx+2)
	}

	return root
}
