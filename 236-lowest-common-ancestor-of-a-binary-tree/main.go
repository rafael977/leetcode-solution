package main

import (
	"fmt"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	} else if l != nil && r == nil {
		return l
	} else if l == nil && r != nil {
		return r
	} else {
		return nil
	}
}

func main() {
	root := BuildTree("3,5,1,6,2,0,8,null,null,7,4")
	fmt.Printf("%#v\n", lowestCommonAncestor(root, NewTreeNode(5), NewTreeNode(1)))

	root = BuildTree("3,5,1,6,2,0,8,null,null,7,4")
	fmt.Printf("%#v\n", lowestCommonAncestor(root, NewTreeNode(5), NewTreeNode(4)))

	root = BuildTree("1,2")
	fmt.Printf("%#v\n", lowestCommonAncestor(root, NewTreeNode(1), NewTreeNode(2)))
}
