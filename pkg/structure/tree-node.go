package structure

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

// BuildTree build  binary tree
func BuildTree(input string) *TreeNode {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}

	var parts = strings.Split(input, ",")
	item := parts[0]
	itemVal, _ := strconv.Atoi(parts[0])
	root := &TreeNode{
		Val:   itemVal,
		Left:  nil,
		Right: nil,
	}
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)

	i := 1
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if i == len(parts) {
			break
		}

		item = strings.TrimSpace(parts[i])
		i++

		if item != "null" {
			leftNumber, _ := strconv.Atoi(item)
			node.Left = &TreeNode{
				Val:   leftNumber,
				Left:  nil,
				Right: nil,
			}
			nodeQueue = append(nodeQueue, node.Left)
		}

		if i == len(parts) {
			break
		}

		item = strings.TrimSpace(parts[i])
		i++
		if item != "null" {
			rightNumber, _ := strconv.Atoi(item)
			node.Right = &TreeNode{
				Val:   rightNumber,
				Left:  nil,
				Right: nil,
			}
			nodeQueue = append(nodeQueue, node.Right)
		}
	}

	return root
}

// PrintTree prints  tree
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val)

	PrintTree(root.Left)
	PrintTree(root.Right)
}
