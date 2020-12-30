package main

import (
	. "github.com/rafael977/leetcode-solution/datastruct"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	cur := root
	done := false
	for !done {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = NewTreeNode(val)
				done = true
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = NewTreeNode(val)
				done = true
			} else {
				cur = cur.Right
			}
		}
	}

	return root
}

func main() {
	root := BuildTree("4,2,7,1,3")
	root = insertIntoBST(root, 5)
	PrintTree(root)

	root = BuildTree("40,20,60,10,30,50,70")
	root = insertIntoBST(root, 25)
	PrintTree(root)

	root = BuildTree("4,2,7,1,3,null,null,null,null,null,null")
	root = insertIntoBST(root, 5)
	PrintTree(root)
}
