package datastruct

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

// NewTreeNode creates new TreeNode
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

// BuildTree build  binary tree
func BuildTree(input string) *TreeNode {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil
	}

	var parts = strings.Split(input, ",")
	item := parts[0]
	itemVal, _ := strconv.Atoi(item)
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
func PrintTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	output := "["
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if node == nil {
			output = fmt.Sprintf("%snull, ", output)
			continue
		}

		output = fmt.Sprintf("%s%d, ", output, node.Val)
		nodeQueue = append(nodeQueue, node.Left)
		nodeQueue = append(nodeQueue, node.Right)
	}
	return output[:len(output)-2] + "]"
}
