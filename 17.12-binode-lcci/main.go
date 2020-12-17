package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode defines for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var head *TreeNode = &TreeNode{-1, nil, nil}
var prev *TreeNode = nil

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)

	root.Left = nil
	if prev == nil {
		head.Right = root
	} else {
		prev.Right = root
	}
	prev = root

	dfs(root.Right)
}

func convertBiNode(root *TreeNode) *TreeNode {
	dfs(root)
	return head.Right
}

func main() {
	arr := processInput("4,2,5,1,3,null,6,0")
	root := buildTree(arr, 0)
	inOrder(root)

	fmt.Println()
	fmt.Println("Converted")
	cov := convertBiNode(root)
	inOrder(cov)
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

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val)
	inOrder(root.Right)
}
